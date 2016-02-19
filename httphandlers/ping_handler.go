package httphandlers

import (
	"fmt"
	"net/http"
)

//Ping is a simple handler that always responds with pong as text
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
