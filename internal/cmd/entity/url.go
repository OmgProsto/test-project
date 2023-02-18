package entity

type Url struct {
	Url          string
	Checks       []string
	MinChecksCnt int `json:"min_checks_cnt"`
}
