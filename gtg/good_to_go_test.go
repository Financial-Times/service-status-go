package g2g

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCreateEmptyStatusChecks(t *testing.T) {
	assert := assert.New(t)
	checkers := GoodToGoChecks{}
	assert.Zero(checkers)
}

func TestCanRunOneStatusCheckThatAlwaysFails(t *testing.T) {
	assert := assert.New(t)
	statusCheck := GoodToGoChecks{
		Checkers: []StatusChecker{localError},
	}
	status := statusCheck.RunChecks()
	assert.False(status.GoodToGo)
	assert.EqualValues(status.Message, localErrorMessage)
}

func TestWillRunOneOfManyStatusCheckThatAlwaysFail(t *testing.T) {
	assert := assert.New(t)
	statusCheck := GoodToGoChecks{
		Checkers: []StatusChecker{localError, localError, localError},
	}
	status := statusCheck.RunChecks()
	assert.False(status.GoodToGo)
	assert.EqualValues(localErrorMessage, status.Message)
}

func TestWillRunAllOfManyStatusCheckThatAlwaysFail(t *testing.T) {
	assert := assert.New(t)
	statusCheck := GoodToGoChecks{
		Checkers:     []StatusChecker{localError, localError, localError},
		RunAllChecks: true,
	}
	status := statusCheck.RunChecks()
	assert.False(status.GoodToGo)
	assert.EqualValues(localErrorMessage+"\n"+localErrorMessage+"\n"+localErrorMessage+"\n", status.Message)
}

func TestCanRunOneStatusCheckThatAlwaysPasses(t *testing.T) {
	assert := assert.New(t)
	statusCheck := GoodToGoChecks{
		Checkers: []StatusChecker{localNoError},
	}
	status := statusCheck.RunChecks()
	assert.True(status.GoodToGo)
}

func TestCanRunManyStatusCheckThatAlwaysPass(t *testing.T) {
	assert := assert.New(t)
	statusCheck := GoodToGoChecks{
		Checkers: []StatusChecker{localNoError, localNoError, localNoError},
	}
	status := statusCheck.RunChecks()
	assert.True(status.GoodToGo)
	assert.EqualValues("OK", status.Message)
}

var localErrorMessage = "There is a problem with the wibble setting, please adjust your set"

func localError() (status Status) {
	status.Message = localErrorMessage
	status.GoodToGo = false
	return status
}

func localNoError() (status Status) {
	status.Message = "This is ignored"
	status.GoodToGo = true
	return status
}
