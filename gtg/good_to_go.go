package gtg

// Implementation of the [FT Good To Go standard](https://docs.google.com/document/d/11paOrAIl9eIOqUEERc9XMaaL3zouJDdkmb-gExYFnw0)

import (
	"time"
)

const (
	timeoutMessage = "Timeout running status check"
	timeout        = 3 * time.Second
)

// Status is the result of running a checker, if the service is GoodToGo then it can serve requests.
// If the message isn't GoodToGo then the message should be in plain text "describing the nature of the problem that prevents the application being good to go.
// This text should be sufficient for a non-domain expert to be able to resolve the problem"
type Status struct {
	Message  string
	GoodToGo bool
}

// StatusChecker is the function signature which a checker needs to implement (no parameters and returns a Status)
type StatusChecker func() Status

// FailFastSequentialChecker composes multiple checkers into one that are executed in sequence. Execution stops as soon as on checker fails.
func FailFastSequentialChecker(checkers []StatusChecker) StatusChecker {
	return func() Status {
		for _, checker := range checkers {
			if status := checker.RunCheck(); !status.GoodToGo {
				return status
			}
		}

		return Status{
			GoodToGo: true,
			Message:  "OK",
		}
	}
}

// FailAtEndSequentialChecker composes multiple checkers into one that are executed in sequence. All checkers are executed.
func FailAtEndSequentialChecker(checkers []StatusChecker) StatusChecker {
	return func() Status {
		result := Status{
			GoodToGo: true,
			Message:  "OK",
		}

		for _, checker := range checkers {
			status := checker.RunCheck()
			if !status.GoodToGo {
				result.GoodToGo = false
				if result.Message == "OK" {
					result.Message = status.Message
				} else {
					result.Message += "\n" + status.Message
				}
			}
		}
		return result
	}
}

// FailFastParallelCheck creates a composite checker that will run all checkers simultaneously.
// As soon as any of the checkers fail then the other checkers are ignored.
func FailFastParallelCheck(checkers []StatusChecker) StatusChecker {
	return func() Status {
		statusChannel := make(chan Status, len(checkers))
		for _, checker := range checkers {
			go func(checker StatusChecker) {
				statusChannel <- checker()
			}(checker)
		}
		for range checkers {
			select {
			case status := <-statusChannel:
				if status.GoodToGo == false {
					return status
				}
			}
		}
		return Status{GoodToGo: true}
	}
}

// RunCheck executes a checker and returns the result as a status
func (check StatusChecker) RunCheck() Status {
	statusChannel := make(chan Status, 1)
	go func() {
		status := check()
		if status.GoodToGo {
			status.Message = "OK"
		}
		statusChannel <- status
	}()
	select {
	case status := <-statusChannel:
		return status
	case <-time.After(timeout):
		return Status{
			GoodToGo: false,
			Message:  timeoutMessage,
		}
	}
}
