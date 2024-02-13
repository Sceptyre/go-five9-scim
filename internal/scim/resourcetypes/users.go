package resourcetypes

import (
	"net/http"

	"github.com/Sceptyre/go-five9-scim/internal/mappers"
	"github.com/Sceptyre/go-five9-scim/internal/scim/filter"
	"github.com/Sceptyre/go-five9-scim/internal/sync"

	five9_api "github.com/Sceptyre/go-five9-scim/pkg/five9/api"

	"github.com/elimity-com/scim"
	scim_errors "github.com/elimity-com/scim/errors"
	"github.com/elimity-com/scim/schema"
)

// Implements github.com/elimity-com/scim.ResourceHandler
type userResourceHandler struct{}

func (urh *userResourceHandler) Create(r *http.Request, attributes scim.ResourceAttributes) (scim.Resource, error) {
	// Map data to VCC userInfo object
	userInfo := mappers.MapScimAttributesToFive9UserInfo(&attributes)
	userInfo.GeneralInfo.Id = 1

	// Post to VCC API
	createUserRes, createUserErr := five9_api.CreateUser(userInfo)
	if createUserErr != nil {
		return scim.Resource{}, scim_errors.ScimError{
			Detail: createUserErr.Body.Error.Message,
			Status: 500,
		}
	}

	// Map response userInfo to a scim resource
	scimUser := mappers.MapFive9UserInfoToScimUser(createUserRes.Return)

	// Update sync data
	sync.SyncIdsToUsernames[scimUser.ID] = createUserRes.Return.GeneralInfo.UserName
	sync.SyncData[scimUser.ID] = *scimUser

	return *scimUser, nil
}

func (urh *userResourceHandler) Get(r *http.Request, id string) (scim.Resource, error) {
	// Get username value for ID
	userName, ok := sync.SyncIdsToUsernames[id]
	if !ok {
		return scim.Resource{}, scim_errors.ScimErrorResourceNotFound(id)
	}

	// Get live user info
	userInfoRes, getUserInfoErr := five9_api.GetUserInfo(userName)
	if getUserInfoErr != nil {
		return scim.Resource{}, scim_errors.ScimError{
			Detail: getUserInfoErr.Body.Error.Message,
			Status: 500,
		}
	}

	scimUser := mappers.MapFive9UserInfoToScimUser(userInfoRes.Return)

	return *scimUser, nil
}

func (urh *userResourceHandler) GetAll(r *http.Request, params scim.ListRequestParams) (scim.Page, error) {
	// Get all resources
	filteredUsers := sync.GetMapValues[scim.Resource](sync.SyncData)

	// Process resources through the filter evaluator
	if params.Filter != nil {
		filteredUsers = filter.Filter(params.Filter, filteredUsers)
	}

	// Get output meta data
	totalResults := len(filteredUsers)

	startIndex := params.StartIndex - 1
	finalIndex := len(filteredUsers)

	// Reduce quantity of users to only requested amount
	if len(filteredUsers) > startIndex {
		finalIndex = startIndex + params.Count
		if len(filteredUsers) < finalIndex {
			finalIndex = len(filteredUsers)
		}
	}

	// Output
	return scim.Page{
		Resources:    filteredUsers[startIndex:finalIndex],
		TotalResults: totalResults,
	}, nil
}

func (urh *userResourceHandler) Replace(r *http.Request, id string, attributes scim.ResourceAttributes) (scim.Resource, error) {
	// Get username value for ID
	userName, ok := sync.SyncIdsToUsernames[id]
	if !ok {
		return scim.Resource{}, scim_errors.ScimErrorResourceNotFound(id)
	}

	// Get live user info
	userInfoRes, getUserInfoErr := five9_api.GetUserInfo(userName)
	if getUserInfoErr != nil {
		return scim.Resource{}, scim_errors.ScimError{
			Detail: getUserInfoErr.Body.Error.Message,
			Status: 500,
		}
	}

	// Map data accordingly
	userInfo := mappers.MapScimAttributesToExistingFive9UserInfo(&attributes, &userInfoRes.Return)
	rolesToRemove := mappers.MapFive9UserInfoToRolesToRemove(&userInfo)

	// Perform update request
	modifyUserRes, modifyUserErr := five9_api.ModifyUser(userInfo.GeneralInfo, userInfo.Roles, rolesToRemove)
	if modifyUserErr != nil {
		return scim.Resource{}, scim_errors.ScimError{
			Detail: modifyUserErr.Body.Error.Message,
			Status: 500,
		}
	}

	// Map response to scim resource
	scimUser := mappers.MapFive9UserInfoToScimUser(modifyUserRes.Return)

	// Update sync data
	sync.SyncIdsToUsernames[scimUser.ID] = modifyUserRes.Return.GeneralInfo.UserName
	sync.SyncData[scimUser.ID] = *scimUser

	// Output scim resource
	return *scimUser, nil
}

func (urh *userResourceHandler) Delete(r *http.Request, id string) error {
	// Get username value for ID
	userName, ok := sync.SyncIdsToUsernames[id]
	if !ok {
		return scim_errors.ScimErrorResourceNotFound(id)
	}

	// Get live user info
	userInfoRes, getUserInfoErr := five9_api.GetUserInfo(userName)
	if getUserInfoErr != nil {
		return scim_errors.ScimError{
			Detail: getUserInfoErr.Body.Error.Message,
			Status: 500,
		}
	}

	// Map data accordingly
	userInfo := userInfoRes.Return
	userInfo.GeneralInfo.Active = false

	// Perform update request
	_, modifyUserErr := five9_api.ModifyUser(userInfo.GeneralInfo, userInfo.Roles, []string{})
	if modifyUserErr != nil {
		return scim_errors.ScimError{
			Detail: modifyUserErr.Body.Error.Message,
			Status: 500,
		}
	}

	// Update sync data
	delete(sync.SyncIdsToUsernames, id)
	delete(sync.SyncData, id)

	// Output scim resource
	return nil
}

func (urh *userResourceHandler) Patch(r *http.Request, id string, operations []scim.PatchOperation) (scim.Resource, error) {
	return scim.Resource{}, scim_errors.ScimErrorBadRequest(
		"Unsupported Method",
	)
}

// Schema for filter and mapped attributes
var userSchema = schema.Schema{
	ID: "urn:ietf:params:scim:schemas:core:2.0:User",
	Attributes: []schema.CoreAttribute{
		schema.SimpleCoreAttribute(schema.SimpleBooleanParams(schema.BooleanParams{
			Name:     "active",
			Required: true,
		})),
		schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
			Name:       "userName",
			Required:   true,
			Uniqueness: schema.AttributeUniquenessServer(),
		})),
		schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
			Name:     "email",
			Required: true,
		})),
		schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
			Name:     "displayName",
			Required: true,
		})),
		schema.ComplexCoreAttribute(schema.ComplexParams{
			Name:     "name",
			Required: true,
			SubAttributes: []schema.SimpleParams{
				schema.SimpleStringParams(schema.StringParams{
					Name:     "givenName",
					Required: true,
				}),
				schema.SimpleStringParams(schema.StringParams{
					Name:     "familyName",
					Required: true,
				}),
			},
		}),
		schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
			Name:        "roles",
			Required:    true,
			MultiValued: true,
		})),
	},
}

// Public resource type variable
var UserResourceType = scim.ResourceType{
	Name:     "User",
	Endpoint: "/Users",
	Schema:   userSchema,

	Handler: &userResourceHandler{},
}
