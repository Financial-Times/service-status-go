package httphandlers

import (
	"encoding/json"
	"github.com/Financial-Times/service-status-go/buildinfo"
	"net/http"
)

const (
	// BuildInfoPath follows the FT convention of prefixing metadata with an underscore
	BuildInfoPath = "/__build-info"
	// BuildInfoPathDW follows the DropWizard convention
	BuildInfoPathDW = "/build-info"
)

//BuildInfoHandler provides a JSON representation of the build-info.
func BuildInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	if methodSupported(w, r) {
		if err := json.NewEncoder(w).Encode(buildinfo.GetBuildInfo()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(error(err.Error()))
		}
	}
}
