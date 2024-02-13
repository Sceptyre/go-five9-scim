package api

import (
	"encoding/xml"

	"github.com/Sceptyre/go-five9-scim/pkg/five9"
	"github.com/Sceptyre/go-five9-scim/pkg/five9/models"
)

func CreateUser(userInfo models.Five9UserInfo) (*models.Five9CreateUserResponse, *models.Five9ErrorResponse) {
	requestBody := models.Five9CreateUserRequest{
		Five9RequestEnvelope: models.GetDefaultRequestEnvelope(),
		Body: models.Five9CreateUserRequestBody{
			UserInfo: userInfo,
		},
	}

	b, _ := xml.MarshalIndent(requestBody, "", " ")

	return five9.Request[models.Five9CreateUserResponse]("POST", "https://api.five9.com:443/wsadmin/v13/AdminService", b)

}
