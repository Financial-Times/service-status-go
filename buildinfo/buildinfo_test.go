package buildinfo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Expected string

func SetUp() {
	version = "0.0.1-RC1"
	repository = "https://github.com/Financial-Times/a-ft-service.git"
	revision = "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"
	builder = "go version go1.5.2 darwin/amd64"
	dateTime = "20130605141043"
	Expected = `{"repository":"https://github.com/Financial-Times/a-ft-service.git",` +
		`"version":"0.0.1-RC1",` +
		`"builder":"go version go1.5.2 darwin/amd64",` +
		`"dateTime":"20130605141043",` +
		`"revision":"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"}`
	parseAndConstruct()
}

func TestModel(t *testing.T) {
	SetUp()
	result, err := json.Marshal(GetBuildInfo())
	assert.NoError(t, err)
	assert.JSONEq(t, Expected, string(result))
}

func TestParseRepository(t *testing.T) {
	validGITURLs := []string{"https://github.com/Financial-Times/service-status-go",
		"http://github.com/Financial-Times/service-status-go.git",
		"https://github.com/Financial-Times/service-status-go.git",
		"git@github.com:Financial-Times/service-status-go.git"}
	for _, url := range validGITURLs {
		repository = url
		assert.NoError(t, parseRepository())
	}
	invalidGITURLs := []string{"file:////Users/someone/code/go/src/github.com/Financial-Times/public-brands-api",
		"ftp://host/somePath/service-status-go.git",
		"git@github.com/service-status-go.git"}
	for _, url := range invalidGITURLs {
		repository = url
		assert.Error(t, parseRepository())
	}
}
