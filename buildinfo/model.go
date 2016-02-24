package buildinfo

import (
	"fmt"
	semver "github.com/hashicorp/go-version"
	"regexp"
)

var version string
var repository string
var revision string
var builder string
var dateTime string

// BuildInfo structure
type BuildInfo struct {
	Version    string `json:"version"`
	Repository string `json:"repository"`
	Revision   string `json:"revision"`
	Builder    string `json:"builder"`
	DateTime   string `json:"dateTime"`
}

// GetBuildInfo returns the current buildInfo as set by the ldflags
func GetBuildInfo() BuildInfo {

	if err := parseRepository(); err != nil {
		repository = err.Error()
	}
	if err := parseRevision(); err != nil {
		revision = err.Error()
	}
	if err := parseVersion(); err != nil {
		version = err.Error()
	}
	if err := parseDateTime(); err != nil {
		dateTime = err.Error()
	}

	return BuildInfo{version, repository, revision, builder, dateTime}
}

// currently suport https repositories
const repositorylMatch = "((git|ssh|http(s)?)|(git@[\\w\\.]+))(:(//)?)([\\w\\.@\\:/\\-~]+)(\\.git)(/)?"

//const repositorylMatch = "^(https?:\\/\\/)?([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\/\\w \\.-]*)*\\/?$"

// currently needs to be a sha1 (ala git)
const revisionMatch = "^[0-9a-f]{5,40}$"

// variant of the iso-8601 standard (i.e. without the seperators)
const dateTimeMatch = "^[0-9]{14}"

var repositoryRegex = regexp.MustCompile(repositorylMatch)
var revisionRegex = regexp.MustCompile(revisionMatch)
var dateTimeRegex = regexp.MustCompile(dateTimeMatch)

func parseRepository() error {
	if !repositoryRegex.MatchString(repository) {
		return fmt.Errorf("Repository %s does not match regex %s", repository, repositorylMatch)
	}
	return nil
}

func parseRevision() error {
	if !revisionRegex.MatchString(revision) {
		return fmt.Errorf("Revision %s does not match regex %s", revision, revisionMatch)
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
