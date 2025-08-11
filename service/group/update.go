package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Updater struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *Updater) Update(ctx *types.ServiceContext, req *model.Group) (rsp *model.Group, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group update")
	return rsp, nil
}

func (g *Updater) UpdateBefore(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group update before")
	return nil
}

func (g *Updater) UpdateAfter(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group update after")
	return nil
}
