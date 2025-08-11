package service

import (
	"helloworld/service/group"
	"helloworld/service/user"

	"github.com/forbearing/golib/service"
)

func Init() error {
	service.Register[*group.Creator]()
	service.Register[*group.Deleter]()
	service.Register[*group.Updater]()
	service.Register[*group.Patcher]()
	service.Register[*group.Lister]()
	service.Register[*group.Getter]()
	service.Register[*group.ManyCreator]()
	service.Register[*group.ManyDeleter]()
	service.Register[*group.ManyUpdater]()
	service.Register[*group.ManyPatcher]()
	service.Register[*user.Creator]()
	service.Register[*user.Deleter]()
	service.Register[*user.Updater]()
	service.Register[*user.Lister]()
	service.Register[*user.Getter]()
	service.Register[*user.ManyCreator]()
	service.Register[*user.ManyDeleter]()
	service.Register[*user.ManyUpdater]()
	service.Register[*user.ManyPatcher]()
	return nil
}
