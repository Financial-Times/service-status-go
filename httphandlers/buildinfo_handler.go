package httphandlers

import (
	"encoding/json"
	"github.com/Financial-Times/service-status-go/buildinfo"
	"net/http"
)

type FtHandler func(http.ResponseWriter, *http.Request)

//BuildInfoHandler is a HandlerFunc that returns a JSON representation of the build-info.
func BuildInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(buildinfo.GetBuildInfo()); err != nil {
		panic(err)
	}
}

func (f FtHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		f(w, r)
	} else {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
