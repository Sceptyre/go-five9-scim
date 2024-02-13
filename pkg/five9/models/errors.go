package models

// Errors
type Five9ErrorResponse struct {
	Env string `xml:"xmlns:env,attr"`

	Body struct {
		Error struct {
			Code    string   `xml:"faultcode"`
			Message string   `xml:"faultstring"`
			Detail  struct{} `xml:"detail"`
		} `xml:"Fault"`
	} `xml:"Body"`

	XMLName struct{} `xml:"Envelope"`
	error
}
