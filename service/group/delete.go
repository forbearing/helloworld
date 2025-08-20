package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Deleter struct {
	service.Base[*model.Group, *model.GroupReq, *model.GroupRsp]
}

func (g *Deleter) Delete(ctx *types.ServiceContext, req *model.GroupReq) (rsp *model.GroupRsp, err error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group delete")
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
