package entity

type Urls struct {
	Urls []Site
}

func (u *Urls) LenUrls() int {
	return len(u.Urls)
}
