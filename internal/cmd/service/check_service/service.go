package check_service

type CheckService struct {
}

func (c CheckService) Check(check Check) (bool, []string) {

	var invalidCheck []string

	for _, v := range check.Checks {
		switch v {
		case "status_code":
			if !c.checkStatusCode(check.Data.StatusCode) {
				invalidCheck = append(invalidCheck, v)
			}

		case "text":
			if !c.checkText(check.Data.Text) {
				invalidCheck = append(invalidCheck, v)
			}
		}
	}

	return len(invalidCheck) < check.MinChecksCount, invalidCheck
}

func (c CheckService) checkStatusCode(statusCode int) bool {
	return statusCode == 200
}

func (c CheckService) checkText(text string) bool {
	return text == "ok"
}
