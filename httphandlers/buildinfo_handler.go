package httphandlers

import (
	"encoding/json"
	"github.com/Financial-Times/service-status-go/buildinfo"
	"net/http"
)

// BuildInfoHandler uses the buildinfo.GetBuildInfo to respond to requests for build info !
func BuildInfoHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(buildinfo.GetBuildInfo()); err != nil {
		panic(err)
	}
}
