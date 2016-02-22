package buildinfo

import (
	"encoding/json"
	"net/http"
)

func BuildInfoHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(GetBuildInfo()); err != nil {
		panic(err)
	}
}
