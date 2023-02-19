package main

import (
	"testproject/internal/cmd/handlers/check_site"
)

func main() {
	cHandler := check_site.NewCheckSiteHandler()

	cHandler.Handle()
}
