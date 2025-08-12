package router

import (
	"helloworld/model"

	"github.com/forbearing/golib/router"
	"github.com/forbearing/golib/types/consts"
)

func Init() error {
	router.Register[*model.Group, *model.Group, *model.Group](router.API(), "group", consts.Get)
	router.Register[*model.Group, *model.Group, *model.Group](router.API(), "group", consts.CreateMany)
	return nil
}
