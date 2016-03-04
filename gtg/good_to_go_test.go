package g2g

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCreateEmptyStatus(t *testing.T) {
	assertions := assert.New(t)
	g2gStatus := Status{}
	assertions.Empty(g2gStatus)
}

func TestCanCreateStatus(t *testing.T) {

}

func localServiceError() error {
	return ServiceError
}
