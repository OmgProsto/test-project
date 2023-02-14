package entity

type Site struct {
	url            string
	checks         []Check
	minChecksCount int
}

func NewSite(url string, checks []Check, minChecksCount int) Site {
	return Site{
		url,
		checks,
		minChecksCount,
	}
}

func (s Site) GetUrl() string {
	return s.url
}

func (s Site) GetChecks() []Check {
	return s.checks
}

func (s Site) GetMinChecksCount() int {
	return s.minChecksCount
}
