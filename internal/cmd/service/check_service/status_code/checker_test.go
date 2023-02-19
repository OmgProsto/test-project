package status_code

import "testing"

func TestStatusCodeChecker_Check200(t *testing.T) {
	statusCodeChecker := StatusCodeChecker{
		statusCode: 200,
	}

	assertCheck(t, statusCodeChecker, true)
}

func TestStatusCodeChecker_CheckNot200(t *testing.T) {
	statusCodeChecker := StatusCodeChecker{
		statusCode: 500,
	}

	assertCheck(t, statusCodeChecker, false)
}

func assertCheck(t *testing.T, checker StatusCodeChecker, want bool) {
	if checker.Check() != want {
		t.Errorf("got %t want %t", checker.Check(), want)
	}
}
