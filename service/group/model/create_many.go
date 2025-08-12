package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type ManyCreator struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *ManyCreator) CreateMany(ctx *types.ServiceContext, req *model.Group) (rsp *model.Group, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group create many")
	return rsp, nil
}

func (g *ManyCreator) CreateManyBefore(ctx *types.ServiceContext, groups ...*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group create many before")
	return nil
}

func (g *ManyCreator) CreateManyAfter(ctx *types.ServiceContext, groups ...*model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group create many after")
	return nil
}
