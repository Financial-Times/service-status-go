package g2g

// Implementation of the [FT Good To Go standard](https://docs.google.com/document/d/11paOrAIl9eIOqUEERc9XMaaL3zouJDdkmb-gExYFnw0)

// Status is the result of running a checker, if the service is GoodToGo then it can serve requests. If the message isn't GoodToGo then the message should be in plain text "describing the nature of the problem that prevents the application being good to go.  This text should be sufficient for a non-domain expert to be able to resolve the problem"
type Status struct {
	Message  string
	GoodToGo bool
}

// StatusChecker is the function signature which a checker needs to implement (no parameters and returns a Status)
type StatusChecker func() Status

// FailFastSequentialChecker composes multiple checkers into one that are executed in sequence. Execution stops as soon as on checker fails.
func FailFastSequentialChecker(checkers []StatusChecker) (checker StatusChecker) {
	f := func() Status {
		for i := range checkers {
			status := checkers[i].RunCheck()
			if !status.GoodToGo {
				return status
			}
		}
		status := Status{
			GoodToGo: true,
			Message:  "OK",
		}
		return status
	}
	return f
}

// FailAtEndSequentialChecker composes multiple checkers into one that are executed in sequence. All checkers are executed.
func FailAtEndSequentialChecker(checkers []StatusChecker) (checker StatusChecker) {
	f := func() Status {
		result := Status{
			GoodToGo: true,
			Message:  "OK",
		}
		for i := range checkers {
			status := checkers[i].RunCheck()
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
	return f
}

// RunCheck executes a checker
func (check StatusChecker) RunCheck() Status {
	status := check()
	if status.GoodToGo {
		status.Message = "OK"
	}
	return status
}
