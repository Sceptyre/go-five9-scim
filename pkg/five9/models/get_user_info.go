package models

// Get User Info Request
type Five9GetUserInfoRequest struct {
	Five9RequestEnvelope
	Body Five9GetUserInfoRequestBody `xml:"soapenv:Body"`
}
type Five9GetUserInfoRequestBody struct {
	UserName string `xml:"ser:getUserInfo>userName"`
}

// Get Users Info
type Five9GetUsersInfoRequest struct {
	Five9RequestEnvelope
	Body Five9GetUsersInfoRequestBody `xml:"soapenv:Body"`
}
type Five9GetUsersInfoRequestBody struct {
	UserNamePattern string `xml:"ser:getUsersInfo>userNamePattern"`
}

// Get User Info Response
type Five9GetUserInfoResponse struct {
	Return  Five9UserInfo `xml:"Body>getUserInfoResponse>return"`
	XMLName struct{}      `xml:"Envelope"`
}

// Get Users Info Response
type Five9GetUsersInfoResponse struct {
	Return  []Five9UserInfo `xml:"Body>getUsersInfoResponse>return"`
	XMLName struct{}        `xml:"Envelope"`
}
