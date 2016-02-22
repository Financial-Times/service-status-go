package buildinfo

import (
	"fmt"
	"time"
)

// Version (x.y.z) is the [semver](http://semver.org) of this build
var Version string

// Repository is where the code was located, eg. the github url
var Repository string

// Commit or revision of the source tree that was build (eg. the sha1 hash from `gitgit rev-parse HEAD`)
var Commit string

// Builder is more freeform, should include the the signature of the thing that built the code, eg. output of 'go version'
var Builder string

var now = time.Now()

var dateTime = fmt.Sprintf("%d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())

// BuildInfo structure to return (model)
type buildInfo struct {
	Repository string `json:"repository"`
	Commit     string `json:"commit"`
	Version    string `json:"version"`
	Builder    string `json:"builder"`
	DateTime   string `json:"dateTime"`
}

func GetBuildInfo() buildInfo {
	return buildInfo{Repository, Commit, Version, Builder, dateTime}
}
