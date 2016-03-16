package httphandlers

import (
	"fmt"
	"net/http"
)

const (
	methodNotAllowed = "Method %s not allowed"
	contentType      = "Content-Type"
	applicationJSON  = "application/json; charset=UTF-8"
	plainText        = "text/plain; charset=US-ASCII"
	cacheControl     = "Cache-control"
	noCache          = "no-cache"
)

func error(msg string) []byte {
	return []byte(fmt.Sprintf(`{"error":"%s"}`, msg))
}

func methodSupported(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == "GET" || r.Method == "HEAD" {
		return true
	}
	w.Header().Set("Allow", "GET, HEAD")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(error(fmt.Sprintf(methodNotAllowed, r.Method)))
	return false
}
