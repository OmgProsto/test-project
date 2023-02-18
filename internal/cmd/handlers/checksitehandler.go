package handlers

import (
	"sync"
	"testproject/internal/cmd/entity"
	"testproject/internal/cmd/service"
)

type CheckSiteHandler struct {
	fileParser        service.JsonFileParser
	jsonToUrls        service.JsonToUrlsService
	getDataSite       service.GetDataSiteService
	checker           service.CheckService
	printResultChecks service.PrintResultChecksService
}

func NewCheckSiteHandler() CheckSiteHandler {
	return CheckSiteHandler{
		service.JsonFileParser{},
		service.JsonToUrlsService{},
		service.NewGetDataSiteService(),
		service.CheckService{},
		service.PrintResultChecksService{},
	}
}

func (c CheckSiteHandler) Handle() {
	jsonString := c.fileParser.Parse("../../storage/sites.json")

	urls := c.jsonToUrls.JsonToEntity(jsonString)

	wg := sync.WaitGroup{}

	wg.Add(len(urls.Urls))

	for _, v := range urls.Urls {
		go func(v entity.Url) {
			defer wg.Done()
			dataSite := c.getDataSite.GetDataSite(v.Url)
			ok, invalidChecks := c.checker.Check(dataSite, v.Checks, v.MinChecksCnt)
			c.printResultChecks.Print(v.Url, ok, invalidChecks)
		}(v)
	}

	wg.Wait()
}
