package httphandlers

import (
	"github.com/Financial-Times/service-status-go/gtg"
	"net/http"
)

const (
	// GTGPath follows the FT convention of prefixing metadata with an underscore
	GTGPath = "/__gtg"
)

type goodToGoChecker struct {
	g2g.StatusChecker
}

func NewGoodToGoHandler(checker g2g.StatusChecker) *goodToGoChecker {
	return &goodToGoChecker{checker}

}

// GoodToGoHandler runs the status checks and sends an HTTP status message
func (checker goodToGoChecker) GoodToGoHandler(w http.ResponseWriter, r *http.Request) {
	if methodSupported(w, r) {
		w.Header().Set(contentType, plainText)
		w.Header().Set(cacheControl, noCache)
		status := checker.RunCheck()
		if status.GoodToGo {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(status.Message))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(status.Message))
		}
	}
}
