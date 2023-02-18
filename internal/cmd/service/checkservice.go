package service

import "testproject/internal/cmd/entity/sitedata"

type CheckService struct {
}

func (c CheckService) Check(dataSite sitedata.DataSiteEntity, check []string, count int) (bool, []string) {

	var invalidCheck []string

	for _, v := range check {
		switch v {
		case "status_code":
			if !c.checkStatusCode(dataSite.StatusCode) {
				invalidCheck = append(invalidCheck, v)
			}

		case "text":
			if !c.checkText(dataSite.Text) {
				invalidCheck = append(invalidCheck, v)
			}
		}
	}

	return len(invalidCheck) < count, invalidCheck
}

func (c CheckService) checkStatusCode(statusCode int) bool {
	return statusCode == 200
}

func (c CheckService) checkText(text string) bool {
	return text == "ok"
}
