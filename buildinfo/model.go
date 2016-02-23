package buildinfo

import (
	"fmt"
	semver "github.com/hashicorp/go-version"
	"regexp"
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

	if err := parseRepository(); err != nil {
		repository = err.Error()
	}
	if err := parseCommit(); err != nil {
		commit = err.Error()
	}
	if err := parseVersion(); err != nil {
		version = err.Error()
	}
	if err := parseDateTime(); err != nil {
		dateTime = err.Error()
	}

	return BuildInfo{repository, commit, version, builder, dateTime}
}

// currently suport https repositories
const urlMatch = "^(https?:\\/\\/)?([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\/\\w \\.-]*)*\\/?$"

// currently needs to be a sha1 (ala git)
const commitMatch = "^[0-9a-f]{5,40}$"

// variant of the iso-8601 standard (i.e. without the seperators)
const dateTimeMatch = "^[0-9]{14}"

var urlRegex = regexp.MustCompile(urlMatch)
var commitRegex = regexp.MustCompile(commitMatch)
var dateTimeRegex = regexp.MustCompile(dateTimeMatch)

func parseRepository() error {
	if !urlRegex.MatchString(repository) {
		return fmt.Errorf("Repository %s does not match regex %s", repository, urlMatch)
	}
	return nil
}

func parseCommit() error {
	if !commitRegex.MatchString(commit) {
		return fmt.Errorf("Commit %s does not match regex %s", commit, commitMatch)
	}
	return nil
}

func parseVersion() error {
	_, err := semver.NewVersion(version)
	return err
}

func parseDateTime() error {
	if !dateTimeRegex.MatchString(dateTime) {
		return fmt.Errorf("dateTime %s does not match regex %s", dateTime, dateTimeRegex)
	}
	return nil
}
