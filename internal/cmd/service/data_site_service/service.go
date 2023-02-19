package data_site_service

import (
	"log"
	"testproject/pkg/client/get"
)

type DataSiteService struct {
}

func (g DataSiteService) GetDataSite(url string) DataSiteEntity {
	request := get.Request{}

	resp, err := request.Get(url)

	var bodyString []byte
	_, err = resp.Body.Read(bodyString)

	if err != nil {
		log.Println(err)
	}

	return DataSiteEntity{
		StatusCode: resp.StatusCode,
		Text:       string(bodyString),
	}
}
