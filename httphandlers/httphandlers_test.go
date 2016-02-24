package httphandlers

import (
	//"github.com/Financial-Times/service-status-go/buildinfo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "ping", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	Ping(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, "pong", w.Body.String())
}

func TestBuildInfoHandler(t *testing.T) {
	// buildinfo.SetUp()
	// req, err := http.NewRequest("GET", "build-info", nil)
	// assert.NoError(t, err)
	// w := httptest.NewRecorder()
	// BuildInfoHandler(w, req)
	// assert.Equal(t, w.Code, http.StatusOK)
	// assert.JSONEq(t, buildinfo.Expected, w.Body.String())
}
