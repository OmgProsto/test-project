package entity

type Check struct {
	check string
}

func NewCheck(check string) Check {
	return Check{
		check: check,
	}
}

func (c Check) Get() string {
	return c.check
}
