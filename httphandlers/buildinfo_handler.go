package httphandlers

import (
	"encoding/json"
	"github.com/Financial-Times/service-status-go/buildinfo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

//BuildInfoHandlerFunc is a HandlerFunc that returns a JSON representation of the build-info.
func BuildInfoHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(buildinfo.GetBuildInfo()); err != nil {
		panic(err)
	}
}

//AppendBuildInfoHandler appends the BuildInfoHandler to a gorilla/mux Router.
//It enforces to respond to HTTP GETs only.
func AppendBuildInfoHandler(r *mux.Router) {
	r.Path("/build-info").Handler(handlers.MethodHandler{"GET": http.HandlerFunc(BuildInfoHandlerFunc)})
}
