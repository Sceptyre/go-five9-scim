package httpserver

import (
	"log"
	"net/http"

	"github.com/Sceptyre/go-five9-scim/internal/httpserver/middleware"
	"github.com/Sceptyre/go-five9-scim/internal/scim"
)

func StartServer() {
	http_server := &http.Server{
		Handler: middleware.WithLogging(scim.ScimServer),
		Addr:    ":8080",
	}

	log.Fatal(http_server.ListenAndServe())
}
