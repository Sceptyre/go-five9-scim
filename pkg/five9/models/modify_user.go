package models

// Get User Info Request
type Five9ModifyUserRequest struct {
	Five9RequestEnvelope
	Body Five9ModifyUserRequestBody `xml:"soapenv:Body"`
}

type Five9ModifyUserRequestBody struct {
	UserGeneralInfo Five9UserGeneralInfo `xml:"ser:modifyUser>userGeneralInfo"`
	RolesToSet      Five9UserRoles       `xml:"ser:modifyUser>rolesToSet"`
	RolesToRemove   []string             `xml:"ser:modifyUser>rolesToRemove"`
}

type Five9ModifyUserResponse struct {
	Return  Five9UserInfo `xml:"Body>modifyUserResponse>return"`
	XMLName struct{}      `xml:"Envelope"`
}
