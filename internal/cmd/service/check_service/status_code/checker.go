package status_code

type StatusCodeChecker struct {
	statusCode int
}

func NewStatusCodeChecker(statusCode int) *StatusCodeChecker {
	return &StatusCodeChecker{
		statusCode: statusCode,
	}
}

func (s StatusCodeChecker) Check() bool {
	return s.statusCode == 200
}
