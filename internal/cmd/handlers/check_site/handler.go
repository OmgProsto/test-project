package check_site

import (
	"sync"
	"testproject/internal/cmd/service/check_service"
	"testproject/internal/cmd/service/data_site_service"
	"testproject/internal/cmd/service/print_service"
	"testproject/pkg/json/json_service"
	"testproject/pkg/json/json_service/entity"
)

type CheckSiteHandler struct {
	parser          json_service.FileParser
	urlModel        json_service.UrlModel
	dataSiteService data_site_service.DataSiteService
	checkService    check_service.CheckService
	printService    print_service.PrintResultService
}

func NewCheckSiteHandler() CheckSiteHandler {
	return CheckSiteHandler{
		json_service.FileParser{},
		json_service.UrlModel{},
		data_site_service.DataSiteService{},
		check_service.CheckService{},
		print_service.PrintResultService{},
	}
}

func (c CheckSiteHandler) Handle() {
	jsonString := c.parser.Parse("../../storage/sites.json")

	urls := c.urlModel.JsonToEntity(jsonString)

	wg := sync.WaitGroup{}

	wg.Add(len(urls.Urls))

	for _, v := range urls.Urls {
		go func(v entity.Url) {
			defer wg.Done()
			dataSite := c.dataSiteService.GetDataSite(v.Url)

			checks := check_service.Check{
				Data:           dataSite,
				Checks:         v.Checks,
				MinChecksCount: v.MinChecksCnt,
			}

			ok, invalidChecks := c.checkService.Check(checks)
			c.printService.Print(v.Url, ok, invalidChecks)
		}(v)
	}

	wg.Wait()
}
