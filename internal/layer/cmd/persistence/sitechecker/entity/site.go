package entity

type Site struct {
	Url          string
	Checks       []string
	MinChecksCnt int `json:"min_checks_cnt"`
}

func (s *Site) LenChecks() int {
	return len(s.Checks)
}
