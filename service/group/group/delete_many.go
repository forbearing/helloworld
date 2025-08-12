package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type ManyDeleter struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *ManyDeleter) DeleteMany(ctx *types.ServiceContext, req *model.Group) (rsp *model.Group, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group delete many")
	return rsp, nil
}

func (g *ManyDeleter) DeleteManyBefore(ctx *types.ServiceContext, groups ...*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group delete many before")
	return nil
}

func (g *ManyDeleter) DeleteManyAfter(ctx *types.ServiceContext, groups ...*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group delete many after")
	return nil
}
