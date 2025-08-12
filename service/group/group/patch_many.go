package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type ManyPatcher struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *ManyPatcher) PatchMany(ctx *types.ServiceContext, req *model.Group) (rsp *model.Group, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group patch many")
	return rsp, nil
}

func (g *ManyPatcher) PatchManyBefore(ctx *types.ServiceContext, groups ...*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group patch many before")
	return nil
}

func (g *ManyPatcher) PatchManyAfter(ctx *types.ServiceContext, groups ...*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group patch many after")
	return nil
}
