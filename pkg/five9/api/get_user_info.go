package api

import (
	"encoding/xml"

	"github.com/Sceptyre/go-five9-scim/pkg/five9"
	"github.com/Sceptyre/go-five9-scim/pkg/five9/models"
)

func GetUserInfo(userName string) (*models.Five9GetUserInfoResponse, *models.Five9ErrorResponse) {
	requestBody := models.Five9GetUserInfoRequest{
		Five9RequestEnvelope: models.GetDefaultRequestEnvelope(),
		Body: models.Five9GetUserInfoRequestBody{
			UserName: userName,
		},
	}

	b, _ := xml.MarshalIndent(requestBody, "", " ")

	return five9.Request[models.Five9GetUserInfoResponse]("POST", "https://api.five9.com:443/wsadmin/v13/AdminService", b)
}
