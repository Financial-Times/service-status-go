package buildinfo

import (
	"github.com/hashicorp/go-version"
	"regexp"
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

var DateTime string

// BuildInfo structure to return (model)
type buildInfo struct {
	Repository string `json:"repository"`
	Commit     string `json:"commit"`
	Version    string `json:"version"`
	Builder    string `json:"builder"`
	DateTime   string `json:"dateTime"`
}

func GetBuildInfo() buildInfo {

	if Repository != "" {
		checkRepository()
	} else {
		Repository = "unknown"
	}

	if Commit != "" {
		checkCommit()
	} else {
		Commit = "unknown"
	}

	if Version != "" {
		checkVersion()
	} else {
		Version = "unknown"
	}

	if Builder == "" {
		Builder = "unknown"
	}

	if DateTime != "" {
		checkDateTime()
	} else {
		DateTime = "unknown"
	}

	return buildInfo{Repository, Commit, Version, Builder, DateTime}
}

const urlRegex = "^(https?:\\/\\/)?([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\/\\w \\.-]*)*\\/?$"

func checkRepository() {
	if regexp.MustCompile(urlRegex).MatchString(Repository) != true {
		panic("The repository value should be a URL")
	}
}

const sha1Regex = "^[0-9a-f]{5,40}$"

func checkCommit() {
	if regexp.MustCompile(sha1Regex).MatchString(Commit) != true {
		panic("The commit should be SHA1 git hash")
	}
}

func checkVersion() {
	_, err := version.NewVersion(Version)
	if err != nil {
		panic("Version is not complain with SemVer")
	}
}

func checkDateTime() {
	_, err := time.Parse(time.RFC3339Nano, DateTime)
	if err != nil {
		panic(err)
	}
}
