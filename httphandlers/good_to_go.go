package httphandlers

import (
	"github.com/Financial-Times/service-status-go/gtg"
	"net/http"
)

const (
	// GTGPath follows the FT convention of prefixing metadata with an underscore
	GTGPath = "/__gtg"
)

type goodToGoChecks struct {
	g2g.GoodToGoChecks
}

func NewGoodToGoHandler(checkers g2g.GoodToGoChecks) *goodToGoChecks {
	return &goodToGoChecks{checkers}

}

// GoodToGoHandler runs the status checks and sends an HTTP status message
func (checkers goodToGoChecks) GoodToGoHandler(w http.ResponseWriter, r *http.Request) {
	if methodSupported(w, r) {
		w.Header().Set(contentType, plainText)
		w.Header().Set(cacheControl, noCache)
		status := checkers.RunChecks()
		if status.GoodToGo {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(status.Message))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(status.Message))
		}
	}
}
