package api

import (
	"encoding/xml"

	"github.com/Sceptyre/go-five9-scim/pkg/five9"
	"github.com/Sceptyre/go-five9-scim/pkg/five9/models"
)

func GetUsersInfo(userNamePattern string) (*models.Five9GetUsersInfoResponse, *models.Five9ErrorResponse) {
	requestBody := models.Five9GetUsersInfoRequest{
		Five9RequestEnvelope: models.GetDefaultRequestEnvelope(),
		Body: models.Five9GetUsersInfoRequestBody{
			UserNamePattern: userNamePattern,
		},
	}

	b, _ := xml.MarshalIndent(requestBody, "", " ")

	return five9.Request[models.Five9GetUsersInfoResponse]("POST", "https://api.five9.com:443/wsadmin/v13/AdminService", b)
}
