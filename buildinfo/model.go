package buildinfo

import (
	"fmt"
	semver "github.com/hashicorp/go-version"
	"regexp"
	"time"
)

var version string
var repository string
var commit string
var builder string
var dateTime string

// BuildInfo structure
type BuildInfo struct {
	Repository string `json:"repository"`
	Commit     string `json:"commit"`
	Version    string `json:"version"`
	Builder    string `json:"builder"`
	DateTime   string `json:"dateTime"`
}

// GetBuildInfo returns the current buildInfo as set by the ldflags
func GetBuildInfo() BuildInfo {

	if err := parseRepository(repository); err != nil {
		repository = err.Error()
	}

	if commit != "" {
		checkCommit()
	} else {
		commit = "unknown"
	}

	if version != "" {
		checkVersion()
	} else {
		version = "unknown"
	}

	if builder == "" {
		builder = "unknown"
	}

	if dateTime != "" {
		checkDateTime()
	} else {
		dateTime = "unknown"
	}

	return BuildInfo{repository, commit, version, builder, dateTime}
}

const urlMatchRegex = "^(https?:\\/\\/)?([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\/\\w \\.-]*)*\\/?$"

var urlRegex *regexp.Regexp

func parseRepository(repository string) (err error) {
	if urlRegex == nil {
		urlRegex, err = regexp.Compile(urlMatchRegex)
	}
	if (err != nil) && (urlRegex != nil) {
		if !urlRegex.Match([]byte(repository)) {
			err = fmt.Errorf("Repository value %s does not match regex %s", repository, urlMatchRegex)
		}
	}
	return err
}

const sha1Regex = "^[0-9a-f]{5,40}$"

func checkCommit() {
	if regexp.MustCompile(sha1Regex).MatchString(commit) != true {
		panic("The commit should be SHA1 git hash")
	}
}

func checkVersion() {
	_, err := semver.NewVersion(version)
	if err != nil {
		panic("Version is not complain with SemVer")
	}
}

func checkDateTime() {
	_, err := time.Parse(time.RFC3339Nano, dateTime)
	if err != nil {
		panic(err)
	}
}
