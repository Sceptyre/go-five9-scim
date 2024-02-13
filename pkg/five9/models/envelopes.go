package models

import (
	"encoding/xml"
)

// Composition Base For Requests
type Five9RequestEnvelope struct {
	SoapENV string   `xml:"xmlns:soapenv,attr"`
	Ser     string   `xml:"xmlns:ser,attr"`
	Header  struct{} `xml:"soapenv:Header"`
	Body    struct{} `xml:"soapenv:Body"`

	XMLName xml.Name `xml:"soapenv:Envelope"`
}

// Composition Base For Responses
type Five9ResponseEnvelope struct {
	Env    string   `xml:"xmlns:env,attr"`
	Header struct{} `xml:"Header"`
	Body   struct{} `xml:"Body"`

	XMLName xml.Name `xml:"Envelope"`
}

var defaultRequestEnvelope = Five9RequestEnvelope{
	SoapENV: "http://schemas.xmlsoap.org/soap/envelope/",
	Ser:     "http://service.admin.ws.five9.com/",
	Header:  struct{}{},
}

func GetDefaultRequestEnvelope() Five9RequestEnvelope {
	return defaultRequestEnvelope
}
