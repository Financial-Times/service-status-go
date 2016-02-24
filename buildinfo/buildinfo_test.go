package buildinfo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
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

}

func TestModel(t *testing.T) {
	SetUp()
	result, err := json.Marshal(GetBuildInfo())
	assert.NoError(t, err)
	assert.JSONEq(t, Expected, string(result))
}
