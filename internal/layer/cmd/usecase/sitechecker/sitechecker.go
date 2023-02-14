package sitechecker

import (
	"testproject/internal/layer/cmd/domain/sitechecker"
)

type CheckSiteUseCase struct {
	FileParser sitechecker.FileParser
}

func (c *CheckSiteUseCase) Check() {
	c.FileParser.Parse()
}
