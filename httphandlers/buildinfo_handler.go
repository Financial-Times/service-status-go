package httphandlers

import (
	"encoding/json"
	"github.com/Financial-Times/service-status-go/buildinfo"
	"net/http"
)

// FtHandler looks like a standard handler to me
type FtHandler func(http.ResponseWriter, *http.Request)

//BuildInfo provides a JSON representation of the build-info.
func BuildInfo(w http.ResponseWriter, r *http.Request) {
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

type httpMux interface {
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}

// RegisterPing just registeres the Ping handler with an http mux
func RegisterPing(mux httpMux) {
	mux.HandleFunc("/__ping", Ping)
	mux.HandleFunc("/ping", Ping)
}

// RegisterBuildInfo adds the build-info handlers
func RegisterBuildInfo(mux httpMux) {
	mux.HandleFunc("/__build-info", BuildInfo)
	mux.HandleFunc("/build-info", BuildInfo)
}

// RegisterAll adds all the handlers
func RegisterAll(mux httpMux) {
	RegisterPing(mux)
	RegisterBuildInfo(mux)
}
