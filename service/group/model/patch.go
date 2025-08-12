package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Patcher struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *Patcher) Patch(ctx *types.ServiceContext, req *model.Group) (rsp *model.Group, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group patch")
	return rsp, nil
}

func (g *Patcher) PatchBefore(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group patch before")
	return nil
}

func (g *Patcher) PatchAfter(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group patch after")
	return nil
}
