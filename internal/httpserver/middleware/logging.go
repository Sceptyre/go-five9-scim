package middleware

import (
	"fmt"
	"net/http"

	"github.com/Sceptyre/go-five9-scim/internal/logger"
)

var Logger = logger.Logger{
	Namespace: "HTTP_SERVER",
}

func WithLogging(h http.Handler) http.Handler {
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		Logger.Log(fmt.Sprintf("%v >> %s %v", r.RemoteAddr, r.Method, r.RequestURI))
		h.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(logFn)
}
