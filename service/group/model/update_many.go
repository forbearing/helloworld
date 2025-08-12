package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type ManyUpdater struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *ManyUpdater) UpdateMany(ctx *types.ServiceContext, req *model.Group) (rsp *model.Group, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group update many")
	return rsp, nil
}

func (g *ManyUpdater) UpdateManyBefore(ctx *types.ServiceContext, groups ...*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group update many before")
	return nil
}

func (g *ManyUpdater) UpdateManyAfter(ctx *types.ServiceContext, groups ...*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group update many after")
	return nil
}
