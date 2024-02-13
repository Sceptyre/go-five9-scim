package five9

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Sceptyre/go-five9-scim/pkg/five9/models"
)

var client = &http.Client{}
var authString string = ""

func Login(username string, password string) {
	auth := username + ":" + password
	authString = base64.StdEncoding.EncodeToString([]byte(auth))
}

func Request[T any](method string, uri string, body []byte) (*T, *models.Five9ErrorResponse) {
	var responseData T
	var errorData models.Five9ErrorResponse

	req, _ := http.NewRequest(method, uri, strings.NewReader(string(body)))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", authString))

	res, err := client.Do(req)

	responseBytes, _ := io.ReadAll(res.Body)

	if err != nil || res.StatusCode != 200 {
		xml.Unmarshal(responseBytes, &errorData)

		return nil, &errorData
	} else {
		xml.Unmarshal(responseBytes, &responseData)

		return &responseData, nil
	}
}
