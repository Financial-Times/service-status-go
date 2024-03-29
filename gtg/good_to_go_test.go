package gtg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCanRunOneStatusCheckThatAlwaysFails(t *testing.T) {
	assert := assert.New(t)
	status := StatusChecker(localError).RunCheck()
	assert.False(status.GoodToGo)
	assert.EqualValues(status.Message, localErrorMessage)
}

func TestWillRunOneOfManyStatusChecksThatAlwaysFail(t *testing.T) {
	assert := assert.New(t)
	statusCheck := FailFastSequentialChecker([]StatusChecker{localError, localError, localError})
	status := statusCheck.RunCheck()
	assert.False(status.GoodToGo)
	assert.EqualValues(localErrorMessage, status.Message)
}

func TestWillRunAllOfManyStatusChecksThatAlwaysFail(t *testing.T) {
	assert := assert.New(t)
	statusCheck := FailAtEndSequentialChecker([]StatusChecker{localError, localError, localError})
	status := statusCheck.RunCheck()
	assert.False(status.GoodToGo)
	assert.EqualValues(localErrorMessage+"\n"+localErrorMessage+"\n"+localErrorMessage, status.Message)
}

func TestCanRunOneStatusCheckThatAlwaysPasses(t *testing.T) {
	assert := assert.New(t)
	status := StatusChecker(localNoError).RunCheck()
	assert.True(status.GoodToGo)
}

func TestCanRunManyStatusChecksThatAlwaysPass(t *testing.T) {
	assert := assert.New(t)
	statusCheck := FailAtEndSequentialChecker([]StatusChecker{localNoError, localNoError, localNoError})
	status := statusCheck.RunCheck()
	assert.True(status.GoodToGo)
	assert.EqualValues("OK", status.Message)
}

func TestTimeoutRunningManyStatusChecksThatAlwaysPass(t *testing.T) {
	assert := assert.New(t)
	statusCheck := FailAtEndSequentialChecker([]StatusChecker{localNoError, delayedChecker(localNoError, 5), localNoError})
	status := statusCheck.RunCheck()
	assert.False(status.GoodToGo)
	assert.EqualValues(timeoutMessage, status.Message)
}

func TestFailFastParallelCheckFailure(t *testing.T) {
	assert := assert.New(t)
	statusCheck := FailFastParallelCheck([]StatusChecker{delayedChecker(localNoError, 1), delayedChecker(localNoError, 1), delayedChecker(localNoError, 1), delayedChecker(localNoError, 1), delayedChecker(localError, 2)})
	status := statusCheck.RunCheck()
	assert.False(status.GoodToGo)
	assert.EqualValues(localErrorMessage, status.Message)
}

func TestFailFastParallelCheckSuccess(t *testing.T) {
	assert := assert.New(t)
	statusCheck := FailFastParallelCheck([]StatusChecker{delayedChecker(localNoError, 2), delayedChecker(localNoError, 2), delayedChecker(localNoError, 2), delayedChecker(localNoError, 2), delayedChecker(localNoError, 2)})
	status := statusCheck.RunCheck()
	assert.True(status.GoodToGo)
	assert.EqualValues("OK", status.Message)
}

func TestFailFastParallelCheckTimeout(t *testing.T) {
	assert := assert.New(t)
	statusCheck := FailAtEndSequentialChecker([]StatusChecker{delayedChecker(localNoError, 1), delayedChecker(localNoError, 1), delayedChecker(localNoError, 1), delayedChecker(localNoError, 1), delayedChecker(localNoError, 5)})
	status := statusCheck.RunCheck()
	assert.False(status.GoodToGo)
	assert.EqualValues(timeoutMessage, status.Message)
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

func delayedChecker(checker StatusChecker, secondsDelay int) StatusChecker {
	fn := func() Status {
		time.Sleep(time.Second * time.Duration(secondsDelay))
		return checker()
	}
	return fn
}
