package project

import (
	"helloworld/model/setting"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Creator struct {
	service.Base[*setting.Project, *setting.Project, *setting.Project]
}

func (p *Creator) Create(ctx *types.ServiceContext, req *setting.Project) (rsp *setting.Project, err error) {
	log := p.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("project create")
	return rsp, nil
}

func (p *Creator) CreateBefore(ctx *types.ServiceContext, project *setting.Project) error {
	log := p.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("project create before")
	return nil
}

func (p *Creator) CreateAfter(ctx *types.ServiceContext, project *setting.Project) error {
	log := p.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("project create after")
	return nil
}
