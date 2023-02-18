package service

import (
	"testproject/internal/cmd/entity/sitedata"
	"testproject/internal/cmd/repository"
)

type GetDataSiteService struct {
	repo repository.RequestRepository
}

func NewGetDataSiteService() GetDataSiteService {
	return GetDataSiteService{
		repository.RequestRepository{},
	}
}

func (g GetDataSiteService) GetDataSite(url string) sitedata.DataSiteEntity {
	return g.repo.SendRequest(url)
}
