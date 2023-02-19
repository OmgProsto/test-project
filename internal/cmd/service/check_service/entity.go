package check_service

import "testproject/internal/cmd/service/data_site_service"

type Check struct {
	Data           data_site_service.DataSiteEntity
	Checks         []string
	MinChecksCount int
}
