package g2g

// Implementation of the [FT Good To Go standard](https://docs.google.com/document/d/11paOrAIl9eIOqUEERc9XMaaL3zouJDdkmb-gExYFnw0)

import (
	"fmt"
)

type Status struct {
	Message  string
	GoodToGo bool
}

// func (status *Status) Error() string {
// 	return status.Message
// }

type StatusChecker func() Status

// type Settings struct {
// 	RunAllChecks    bool
// 	SerialiseChecks bool
// }

type GoodToGoChecks struct {
	Checkers        []StatusChecker
	RunAllChecks    bool
	SerialiseChecks bool
}

// func NewGoodToGoChecks(checkers []StatusChecker, settings Settings) *GoodToGoChecks {
//         if settings = nil {
//                 settings = Settings{}
//         }
// 	return &GoodToGoChecks{checkers, settings}
// }

func (checks GoodToGoChecks) RunChecks() (overall Status) {
	overall.GoodToGo = true
	for i := range checks.Checkers {
		status := checks.Checkers[i]()
		if !status.GoodToGo {
			overall.GoodToGo = false
			if checks.RunAllChecks {
				overall.Message = fmt.Sprintf("%s%s\n", overall.Message, status.Message)
			} else {
				overall.Message = status.Message
				break
			}
		}
	}
	if overall.GoodToGo {
		overall.Message = "OK"
	}
	return overall
}
