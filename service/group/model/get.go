package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Getter struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *Getter) Get(ctx *types.ServiceContext, req *model.Group) (rsp *model.Group, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group get")
	return rsp, nil
}

func (g *Getter) GetBefore(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group get before")
	return nil
}

func (g *Getter) GetAfter(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group get after")
	return nil
}
