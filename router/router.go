package router

import (
	"helloworld/model"

	"github.com/forbearing/golib/router"
	"github.com/forbearing/golib/types/consts"
)

func Init() error {
	router.Register[*model.User, *model.User, *model.User](router.API(), "user2", consts.Create)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user2", consts.Delete)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user2", consts.Update)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user2", consts.List)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user2", consts.Get)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user2", consts.CreateMany)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user2", consts.DeleteMany)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user2", consts.UpdateMany)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user2", consts.PatchMany)
	return nil
}
