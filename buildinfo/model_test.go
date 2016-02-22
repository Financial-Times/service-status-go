package buildinfo

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefault(t *testing.T) {
	expected := `{"repository":"unknown", "version":"unknown", "builder":"unknown", "dateTime":"` +
		fmt.Sprintf("%d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()) + `", "commit":"unknown"}`
	result, err := json.Marshal(BuildInfo)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(result))
}
