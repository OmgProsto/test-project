package check_service

import (
	"testproject/internal/cmd/service/check_service/status_code"
	"testproject/internal/cmd/service/check_service/text"
)

const checkStatusCode = "status_code"
const checkText = "text"

type CheckService struct {
}

type Checker interface {
	Check() bool
}

func (c CheckService) Check(check Check) (bool, []string) {

	var invalidCheck []string

	for _, v := range check.Checks {
		switch v {
		case checkStatusCode:
			if !c.check(status_code.NewStatusCodeChecker(check.Data.StatusCode)) {
				invalidCheck = append(invalidCheck, v)
			}

		case checkText:
			if !c.check(text.NewTextChecker(check.Data.Text)) {
				invalidCheck = append(invalidCheck, v)
			}
		}
	}

	return len(invalidCheck) < check.MinChecksCount, invalidCheck
}

func (c CheckService) check(checker Checker) bool {
	return checker.Check()
}
