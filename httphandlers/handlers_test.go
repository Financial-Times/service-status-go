package httphandlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "ping", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	PingHandler(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, "pong", w.Body.String())
}

func TestBuildInfoHandler(t *testing.T) {
	expected := `
        {
          "version": "Version  is not a semantic version",
          "repository": "Repository  does not match regex ((git|ssh|http(s)?)|(git@[\\w\\.]+))(:(//)?)([\\w\\.@\\:/\\-~]+)(\\.git)(/)?",
          "revision": "Revision  does not match regex ^[0-9a-f]{5,40}$",
          "builder": "",
          "dateTime": "dateTime  does not match regex ^[0-9]{14}"
        }`
	req, err := http.NewRequest("GET", "build-info", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	BuildInfoHandler(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.JSONEq(t, expected, w.Body.String())
}
