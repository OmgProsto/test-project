package repository

import (
	"log"
	"net/http"
	"testproject/internal/cmd/entity/sitedata"
)

type RequestRepository struct {
}

func (r RequestRepository) SendRequest(url string) sitedata.DataSiteEntity {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	var bodyString []byte
	_, err = resp.Body.Read(bodyString)

	if err != nil {
		log.Println(err)
	}

	return sitedata.DataSiteEntity{
		StatusCode: resp.StatusCode,
		Text:       string(bodyString),
	}
}
