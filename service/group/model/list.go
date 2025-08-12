package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Lister struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *Lister) List(ctx *types.ServiceContext, req *model.Group) (rsp *model.Group, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group list")
	return rsp, nil
}

func (g *Lister) ListBefore(ctx *types.ServiceContext, groups *[]*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group list before")
	return nil
}

func (g *Lister) ListAfter(ctx *types.ServiceContext, groups *[]*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group list after")
	return nil
}
