package router

import (
	"helloworld/model"
	"helloworld/model/cmdb"
	"helloworld/model/setting"

	"github.com/forbearing/golib/router"
	"github.com/forbearing/golib/types/consts"
)

func Init() error {
	router.Register[*cmdb.DNS, *cmdb.DNS, *cmdb.DNS](router.API(), "cmdb/dns", consts.List)
	router.Register[*cmdb.Machine, *cmdb.Machine, *cmdb.Machine](router.API(), "cmdb/machine", consts.List)
	router.Register[*model.Group, *model.GroupReq, *model.GroupRsp](router.API(), "group", consts.Create)
	router.Register[*model.Group, *model.GroupReq, *model.GroupRsp](router.API(), "group", consts.Delete)
	router.Register[*model.Group, *model.GroupReq, *model.GroupRsp](router.API(), "group", consts.Update)
	router.Register[*model.Group, *model.Group, *model.Group](router.API(), "group", consts.Patch)
	router.Register[*model.Group, *model.Group, *model.Group](router.API(), "group", consts.List)
	router.Register[*model.Group, *model.Group, *model.Group](router.API(), "group", consts.Get)
	router.Register[*model.Group, *model.Group, *model.Group](router.API(), "group", consts.CreateMany)
	router.Register[*model.Group, *model.Group, *model.Group](router.API(), "group", consts.DeleteMany)
	router.Register[*model.Group, *model.Group, *model.Group](router.API(), "group", consts.UpdateMany)
	router.Register[*model.Group, *model.Group, *model.Group](router.API(), "group", consts.PatchMany)
	router.Register[*setting.Project, *setting.Project, *setting.Project](router.API(), "setting/project", consts.Create)
	router.Register[*setting.Project, *setting.Project, *setting.Project](router.API(), "setting/project", consts.Delete)
	router.Register[*setting.Project, *setting.Project, *setting.Project](router.API(), "setting/project", consts.Update)
	router.Register[*setting.Region, *setting.Region, *setting.Region](router.API(), "setting/region", consts.Create)
	router.Register[*setting.Region, *setting.Region, *setting.Region](router.API(), "setting/region", consts.List)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user", consts.Create)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user", consts.Delete)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user", consts.Update)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user", consts.List)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user", consts.Get)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user", consts.CreateMany)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user", consts.DeleteMany)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user", consts.UpdateMany)
	router.Register[*model.User, *model.User, *model.User](router.API(), "user", consts.PatchMany)
	return nil
}
