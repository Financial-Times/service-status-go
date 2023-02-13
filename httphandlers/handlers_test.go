package httphandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Financial-Times/service-status-go/gtg"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", PingPath, nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	PingHandler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestBuildInfoHandler(t *testing.T) {
	expected := `
        {
          "version": "Version  is not a semantic version",
          "repository": "Repository  does not match regex ((git|ssh|http(s)?)|(git@[\\w\\.]+))(:(//)?)([\\w\\.@\\:/\\-~]+)(\\.git)?(/)?",
          "revision": "Revision  does not match regex ^[0-9a-f]{5,40}$",
          "builder": "",
          "dateTime": "dateTime  does not match regex ^[0-9]{14}"
        }`
	req, err := http.NewRequest("GET", BuildInfoPath, nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	BuildInfoHandler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}

func TestGTGHandlerNoError(t *testing.T) {
	handler := NewGoodToGoHandler(everythingIsOK)
	req, err := http.NewRequest("GET", GTGPath, nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.EqualValues(t, "OK", w.Body.String())
}

func TestGTGHandlerError(t *testing.T) {
	handler := NewGoodToGoHandler(everythingIsNotOK)
	req, err := http.NewRequest("GET", GTGPath, nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler(w, req)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	assert.NotContains(t, "OK", w.Body.String())
}

func TestGTGHandlerTimeoutError(t *testing.T) {
	handler := NewGoodToGoHandler(takesTooLong)
	req, err := http.NewRequest("GET", GTGPath, nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler(w, req)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	assert.EqualValues(t, "Timeout running status check", w.Body.String())
}

func everythingIsOK() (status gtg.Status) {
	status.GoodToGo = true
	return status
}

func everythingIsNotOK() (status gtg.Status) {
	status.Message = "I'm sorry, Dave. I'm afraid I can't do that."
	status.GoodToGo = false
	return status
}

func takesTooLong() (status gtg.Status) {
	time.Sleep(time.Second * 10)
	status.GoodToGo = true
	return status
}
