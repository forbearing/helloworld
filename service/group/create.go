package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Creator struct {
	service.Base[*model.Group, *model.GroupReq, model.GroupRsp]
}

func (g *Creator) Create(ctx *types.ServiceContext, req *model.GroupReq) (rsp model.GroupRsp, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group create")
	return rsp, nil
}

func (g *Creator) CreateBefore(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group create before")
	return nil
}

func (g *Creator) CreateAfter(ctx *types.ServiceContext, group *model.Group) error {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group create after")
	return nil
}
