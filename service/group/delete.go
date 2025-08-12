package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/database"
	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Deleter struct {
	service.Base[*model.Group, *model.GroupReq, *model.GroupRsp]
}

func (g *Deleter) Delete(ctx *types.ServiceContext, req *model.GroupReq) (rsp *model.GroupRsp, err error) {
	list := make([]*model.Group, 0)
	if err = database.Database[*model.Group]().WithQuery(&model.Group{Name: req.Name}).List(&list); err != nil {
		return nil, err
	}

	if err = database.Database[*model.Group]().Delete(list...); err != nil {
		return nil, err
	}

	rsp = &model.GroupRsp{
		Deleted: list,
	}

	return rsp, nil
}

func (g *Deleter) DeleteBefore(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group delete before")
	return nil
}

func (g *Deleter) DeleteAfter(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group delete after")
	return nil
}
