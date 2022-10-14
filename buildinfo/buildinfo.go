package buildinfo

import (
	"runtime"
	"runtime/debug"
)

func init() {
	parseAndConstruct()
}

// BuildInfo structure
type BuildInfo struct {
	Version    string `json:"version"`
	Repository string `json:"repository"`
	Revision   string `json:"revision"`
	Builder    string `json:"builder"`
	DateTime   string `json:"dateTime"`
}

var buildInfo BuildInfo

// GetBuildInfo returns the current buildInfo as set by the ldflags
func GetBuildInfo() BuildInfo {
	return buildInfo
}

func parseAndConstruct() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		buildInfo = BuildInfo{
			Version:    "Failed to read debug BuildInfo",
			Repository: "",
			Revision:   "",
			Builder:    "",
			DateTime:   "",
		}
		return
	}
	var revision string
	var buildTime string

	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			revision = s.Value
		case "vcs.time":
			buildTime = s.Value
		}
	}
	buildInfo = BuildInfo{
		Version:    info.Main.Version,
		Repository: info.Path,
		Revision:   revision,
		Builder:    runtime.Version(),
		DateTime:   buildTime,
	}
}
