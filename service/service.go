package service

import (
	"helloworld/service/cmdb/dns"
	"helloworld/service/cmdb/machine"
	"helloworld/service/group"
	"helloworld/service/setting/project"
	"helloworld/service/setting/region"
	"helloworld/service/twofa/setup"
	"helloworld/service/twofa/verify"
	"helloworld/service/user"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types/consts"
)

func Init() error {
	service.Register[*dns.Lister](consts.PHASE_LIST)
	service.Register[*machine.Lister](consts.PHASE_LIST)
	service.Register[*group.Creator](consts.PHASE_CREATE)
	service.Register[*group.Deleter](consts.PHASE_DELETE)
	service.Register[*group.Updater](consts.PHASE_UPDATE)
	service.Register[*group.Patcher](consts.PHASE_PATCH)
	service.Register[*group.Lister](consts.PHASE_LIST)
	service.Register[*group.Getter](consts.PHASE_GET)
	service.Register[*group.ManyCreator](consts.PHASE_CREATE_MANY)
	service.Register[*group.ManyDeleter](consts.PHASE_DELETE_MANY)
	service.Register[*group.ManyUpdater](consts.PHASE_UPDATE_MANY)
	service.Register[*group.ManyPatcher](consts.PHASE_PATCH_MANY)
	service.Register[*group.Importer](consts.PHASE_IMPORT)
	service.Register[*group.Exporter](consts.PHASE_EXPORT)
	service.Register[*project.Creator](consts.PHASE_CREATE)
	service.Register[*project.Deleter](consts.PHASE_DELETE)
	service.Register[*project.Updater](consts.PHASE_UPDATE)
	service.Register[*region.Creator](consts.PHASE_CREATE)
	service.Register[*region.Lister](consts.PHASE_LIST)
	service.Register[*setup.Lister](consts.PHASE_LIST)
	service.Register[*verify.Creator](consts.PHASE_CREATE)
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
