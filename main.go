package main

import (
	"fmt"
	"github.com/Financial-Times/service-status-go/buildinfo"
)

func main() {
	info := buildinfo.GetBuildInfo()
	fmt.Printf("Build info is %v+", info)
}
