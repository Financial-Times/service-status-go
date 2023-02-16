package httphandlers

import (
	"github.com/Financial-Times/go-logger/v2"
	"net/http"

	"github.com/Financial-Times/service-status-go/gtg"
)

const (
	// GTGPath follows the FT convention of prefixing metadata with an underscore
	GTGPath = "/__gtg"
)

var log = logger.NewUPPLogger("GTG-Handler", "INFO")

type goodToGoChecker struct {
	gtg.StatusChecker
}

// NewGoodToGoHandler is used to construct a new GoodToGoHandler
func NewGoodToGoHandler(checker gtg.StatusChecker) func(http.ResponseWriter, *http.Request) {
	return goodToGoChecker{checker}.GoodToGoHandler
}

// GoodToGoHandler runs the status checks and sends an HTTP status message
func (checker goodToGoChecker) GoodToGoHandler(w http.ResponseWriter, r *http.Request) {
	if methodSupported(w, r) {
		log.Info("Running GTG handler...")
		w.Header().Set(contentType, plainText)
		w.Header().Set(cacheControl, noCache)
		status := checker.RunCheck()
		if status.GoodToGo {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		log.WithField("status", status).Info("GTG handler finished")
		w.Write([]byte(status.Message))
	}
}
