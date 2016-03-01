package httphandlers

import (
	"fmt"
	"net/http"
)

const (
	// PingPath follows the general FT standards for path as it starts with an underscore, although technically Ping is not a FT Standard healthcheck
	PingPath = "/__ping"
	// PingPathDW is the DropWizard equivalent path, here for compatabillity for monitors that typically talk to DropWizard ms
	PingPathDW = "/ping"
)

// PingHandler is a simple handler that always responds with pong as text
func PingHandler(w http.ResponseWriter, r *http.Request) {
	if methodSupported(w, r) {
		fmt.Fprintf(w, "pong")
	}

}
