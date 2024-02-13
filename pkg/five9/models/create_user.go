package models

// Get User Info Request
type Five9CreateUserRequest struct {
	Five9RequestEnvelope
	Body Five9CreateUserRequestBody `xml:"soapenv:Body"`
}

type Five9CreateUserRequestBody struct {
	UserInfo Five9UserInfo `xml:"ser:createUser>userInfo"`
}

type Five9CreateUserResponse struct {
	Return  Five9UserInfo `xml:"Body>createUserResponse>return"`
	XMLName struct{}      `xml:"Envelope"`
}
