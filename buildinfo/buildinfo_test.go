package buildinfo

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var expected string

func setUp() {
	Version = "0.0.1"
	Repository = "https://github.com/Financial-Times/a-ft-service"
	Commit = "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"
	Builder = "go version go1.5.2 darwin/amd64"
	expected = `{"repository":"https://github.com/Financial-Times/a-ft-service", "version":"0.0.1", "builder":"go version go1.5.2 darwin/amd64", "dateTime":"` +
		fmt.Sprintf("%d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()) + `", "commit":"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"}`

}

func TestModel(t *testing.T) {
	setUp()
	result, err := json.Marshal(GetBuildInfo())
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(result))
}

func TestBuildInfoHandler(t *testing.T) {
	setUp()
	req, err := http.NewRequest("GET", "build-info", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	BuildInfoHandler(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.JSONEq(t, expected, w.Body.String())
}
