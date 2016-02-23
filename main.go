package main

import (
	"fmt"
	"github.com/Financial-Times/service-status-go/buildinfo"
)

func main() {
	fmt.Printf("Build info is %+v", buildinfo.GetBuildInfo)
}
