package model

import (
	domain "testproject/internal/layer/cmd/domain/sitechecker/entity"
	persistence "testproject/internal/layer/cmd/persistence/sitechecker/entity"
)

type SiteModel struct {
}

func (s *SiteModel) ToDomain(urls persistence.Urls) []domain.Site {
	domainSites := make([]domain.Site, urls.LenUrls())

	for i, site := range urls.Urls {
		domainSites[i] = domain.NewSite(site.Url, s.makeDomainCheck(site), site.MinChecksCnt)
	}

	return domainSites
}

func (s *SiteModel) makeDomainCheck(site persistence.Site) []domain.Check {
	domainChecks := make([]domain.Check, site.LenChecks())

	for i, check := range site.Checks {
		domainChecks[i] = domain.NewCheck(check)
	}

	return domainChecks
}
