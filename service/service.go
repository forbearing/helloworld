package service

import (
	"github.com/forbearing/golib/service"
)

func Init() error {
	service.Register[*userCreator]()
	service.Register[*userDeleter]()
	service.Register[*userUpdater]()
	service.Register[*userLister]()
	service.Register[*userGetter]()
	service.Register[*userManyCreator]()
	service.Register[*userManyDeleter]()
	service.Register[*userManyUpdater]()
	service.Register[*userManyPatcher]()
	return nil
}
