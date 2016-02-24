package main

import (
	"encoding/json"
	"fmt"
	"github.com/Financial-Times/service-status-go/buildinfo"
)

func main() {
	if buildInfo, err := json.Marshal(buildinfo.GetBuildInfo()); err == nil {
		fmt.Printf("%s", string(buildInfo))
	}
}
