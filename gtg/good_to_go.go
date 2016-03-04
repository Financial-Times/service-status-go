package g2g

// Implementation of the [FT Good To Go standard](https://docs.google.com/document/d/11paOrAIl9eIOqUEERc9XMaaL3zouJDdkmb-gExYFnw0)

import ()

// A StatusError is returned by the Checker
type StatusError struct {
	Message      string
	WrappedError error
}

func (e *StatusError) Error() string {
	return e.Message
}

// Checked is
type Status struct {
	message string
	checker func() error
}
