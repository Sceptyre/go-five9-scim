package mappers

import (
	"strconv"
	"strings"

	"github.com/Sceptyre/go-five9-scim/internal/pwgen"
	five9_models "github.com/Sceptyre/go-five9-scim/pkg/five9/models"

	"github.com/elimity-com/scim"
)

func mapPermissionsToPermissionStrings(permissions []five9_models.Five9UserPermission, prefix string) []string {
	output := []string{}

	for _, v := range permissions {
		if v.Value {
			output = append(output, prefix+":"+v.Type)
		}
	}

	return output
}

func MapFive9UserInfoToScimUser(f9User five9_models.Five9UserInfo) *scim.Resource {
	scimResourceAttributes := map[string]interface{}{}

	scimResourceAttributes["active"] = f9User.GeneralInfo.Active
	scimResourceAttributes["userName"] = f9User.GeneralInfo.FederationId

	if f9User.GeneralInfo.FederationId == "" {
		scimResourceAttributes["userName"] = f9User.GeneralInfo.UserName
	}

	scimResourceAttributes["name"] = map[string]interface{}{
		"givenName":  f9User.GeneralInfo.FirstName,
		"familyName": f9User.GeneralInfo.LastName,
	}

	scimResourceAttributes["displayName"] = f9User.GeneralInfo.FullName
	scimResourceAttributes["email"] = f9User.GeneralInfo.Email

	rolesList := []string{}

	rolesList = append(rolesList, mapPermissionsToPermissionStrings(f9User.Roles.Admin.Permissions, "admin")...)
	rolesList = append(rolesList, mapPermissionsToPermissionStrings(f9User.Roles.Agent.Permissions, "agent")...)
	rolesList = append(rolesList, mapPermissionsToPermissionStrings(f9User.Roles.Supervisor.Permissions, "supervisor")...)
	rolesList = append(rolesList, mapPermissionsToPermissionStrings(f9User.Roles.Reporting.Permissions, "reporting")...)

	scimResourceAttributes["roles"] = rolesList

	return &scim.Resource{
		ID:         strconv.Itoa(f9User.GeneralInfo.Id),
		Attributes: scimResourceAttributes,
	}
}

func MapFiveUserInfoListToScimUserList(f9Users *[]five9_models.Five9UserInfo) []scim.Resource {
	output := []scim.Resource{}

	for _, v := range *f9Users {
		output = append(output, *MapFive9UserInfoToScimUser(v))
	}

	return output
}

func setIfExists[T any](potentialValue interface{}, currentValue T) T {
	newValue, ok := potentialValue.(T)

	if ok {
		return newValue
	}

	return currentValue
}

func MapScimAttributesToExistingFive9UserInfo(attributes *scim.ResourceAttributes, userInfo *five9_models.Five9UserInfo) five9_models.Five9UserInfo {
	userInfo.GeneralInfo.Email = setIfExists[string](
		(*attributes)["email"],
		userInfo.GeneralInfo.Email,
	)

	userInfo.GeneralInfo.Active = setIfExists[bool](
		(*attributes)["active"],
		userInfo.GeneralInfo.Active,
	)

	userInfo.GeneralInfo.FullName = setIfExists[string](
		(*attributes)["displayName"],
		userInfo.GeneralInfo.FullName,
	)

	userInfo.GeneralInfo.FederationId = setIfExists[string](
		(*attributes)["userName"],
		userInfo.GeneralInfo.FederationId,
	)

	userInfo.GeneralInfo.FirstName = (*attributes)["name"].(map[string]interface{})["givenName"].(string)
	userInfo.GeneralInfo.LastName = (*attributes)["name"].(map[string]interface{})["familyName"].(string)

	for _, permission := range (*attributes)["roles"].([]interface{}) {
		role, perm, _ := strings.Cut(permission.(string), ":")
		switch strings.ToLower(role) {
		case "admin":
			for i, rolePermission := range userInfo.Roles.Admin.Permissions {
				if rolePermission.Type == perm {
					userInfo.Roles.Admin.Permissions[i].Value = true
				}
			}
		case "agent":
			for i, rolePermission := range userInfo.Roles.Agent.Permissions {
				if rolePermission.Type == perm {
					userInfo.Roles.Agent.Permissions[i].Value = true
				}
			}
		case "supervisor":
			for i, rolePermission := range userInfo.Roles.Supervisor.Permissions {
				if rolePermission.Type == perm {
					userInfo.Roles.Supervisor.Permissions[i].Value = true
				}
			}
		case "reporting":
			for i, rolePermission := range userInfo.Roles.Reporting.Permissions {
				if rolePermission.Type == perm {
					userInfo.Roles.Reporting.Permissions[i].Value = true
				}
			}
		}
	}

	userInfo.GeneralInfo.Password = string(pwgen.RandPW())

	return *userInfo
}

func MapScimAttributesToFive9UserInfo(attributes *scim.ResourceAttributes) five9_models.Five9UserInfo {
	userInfo := five9_models.NewFive9UserInfo()
	userInfo.GeneralInfo.UserName = (*attributes)["userName"].(string)

	return MapScimAttributesToExistingFive9UserInfo(attributes, userInfo)
}

func MapFive9UserInfoToRolesToRemove(userInfo *five9_models.Five9UserInfo) []string {
	rolesToRemove := []string{}

	if len(userInfo.Roles.Admin.Permissions) == 0 {
		rolesToRemove = append(rolesToRemove, "admin")
	}
	if len(userInfo.Roles.Supervisor.Permissions) == 0 {
		rolesToRemove = append(rolesToRemove, "supervisor")
	}
	if len(userInfo.Roles.Agent.Permissions) == 0 {
		rolesToRemove = append(rolesToRemove, "agent")
	}
	if len(userInfo.Roles.Reporting.Permissions) == 0 {
		rolesToRemove = append(rolesToRemove, "reporting")
	}

	return rolesToRemove
}
