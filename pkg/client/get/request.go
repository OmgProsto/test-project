package get

import (
	"net/http"
)

type Request struct {
}

func (r Request) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
