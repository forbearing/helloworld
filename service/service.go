package service

import (
	"helloworld/service/group"
	"helloworld/service/user"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types/consts"
)

func Init() error {
	service.Register[*group.Creator](consts.PHASE_CREATE)
	service.Register[*group.Lister](consts.PHASE_LIST)
	service.Register[*group.Getter](consts.PHASE_GET)
	service.Register[*user.Creator](consts.PHASE_CREATE)
	service.Register[*user.Deleter](consts.PHASE_DELETE)
	service.Register[*user.Updater](consts.PHASE_UPDATE)
	service.Register[*user.Lister](consts.PHASE_LIST)
	service.Register[*user.Getter](consts.PHASE_GET)
	service.Register[*user.ManyCreator](consts.PHASE_CREATE_MANY)
	service.Register[*user.ManyDeleter](consts.PHASE_DELETE_MANY)
	service.Register[*user.ManyUpdater](consts.PHASE_UPDATE_MANY)
	service.Register[*user.ManyPatcher](consts.PHASE_PATCH_MANY)
	return nil
}
