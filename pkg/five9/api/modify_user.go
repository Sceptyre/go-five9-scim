package api

import (
	"encoding/xml"

	"github.com/Sceptyre/go-five9-scim/pkg/five9"
	"github.com/Sceptyre/go-five9-scim/pkg/five9/models"
)

func ModifyUser(userGeneralInfo models.Five9UserGeneralInfo, rolesToSet models.Five9UserRoles, rolesToRemove []string) (*models.Five9ModifyUserResponse, *models.Five9ErrorResponse) {
	requestBody := models.Five9ModifyUserRequest{
		Five9RequestEnvelope: models.GetDefaultRequestEnvelope(),
		Body: models.Five9ModifyUserRequestBody{
			UserGeneralInfo: userGeneralInfo,
			RolesToSet:      rolesToSet,
			RolesToRemove:   rolesToRemove,
		},
	}

	b, _ := xml.MarshalIndent(requestBody, "", " ")

	return five9.Request[models.Five9ModifyUserResponse]("POST", "https://api.five9.com:443/wsadmin/v13/AdminService", b)

}
