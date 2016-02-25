package httphandlers

import (
	"encoding/json"
	"github.com/Financial-Times/service-status-go/buildinfo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// BuildInfoHandlerFunc uses the buildinfo.GetBuildInfo to respond to requests for build info !
func BuildInfoHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(buildinfo.GetBuildInfo()); err != nil {
		panic(err)
	}
}

func AppendBuildInfo(r *mux.Router) {
	r.Path("/build-info").Handler(handlers.MethodHandler{"GET": http.HandlerFunc(BuildInfoHandlerFunc)}).Header("Content-Type", "application/json; charset=utf-8")
}
