package main

import "testproject/internal/cmd/handlers"

func main() {
	cHandler := handlers.NewCheckSiteHandler()

	cHandler.Handle()
}
