package service

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
)

func Init() error {
	service.Register[*user]()
	return nil
}

type user struct {
	service.Base[*model.User, *model.User, *model.User]
}
