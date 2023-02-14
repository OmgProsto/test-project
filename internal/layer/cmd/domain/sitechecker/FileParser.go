package sitechecker

import "testproject/internal/layer/cmd/domain/sitechecker/entity"

type FileParser interface {
	Parse() []entity.Site
}
