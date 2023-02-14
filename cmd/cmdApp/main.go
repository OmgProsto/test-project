package main

import (
	"testproject/internal/app/di"
)

func main() {
	container := di.NewContainer()

	container.CheckSiteUseCase().Check()
}
