package router

import (
	"helloworld/model"

	"github.com/forbearing/golib/router"
)

func Init() error {
	router.Register[*model.User, *model.User, *model.User](router.API(), "user")
	return nil
}
