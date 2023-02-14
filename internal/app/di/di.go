package di

import (
	"testproject/internal/layer/cmd/persistence/sitechecker/action"
	sitechecker3 "testproject/internal/layer/cmd/usecase/sitechecker"
)

type Container struct {
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) CheckSiteUseCase() *sitechecker3.CheckSiteUseCase {
	return &sitechecker3.CheckSiteUseCase{
		FileParser: action.FileParserAction{},
	}
}
