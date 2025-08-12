package project

import (
	"helloworld/model/setting"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Updater struct {
	service.Base[*setting.Project, *setting.Project, *setting.Project]
}

func (p *Updater) Update(ctx *types.ServiceContext, req *setting.Project) (rsp *setting.Project, err error) {
	log := p.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("project update")
	return rsp, nil
}

func (p *Updater) UpdateBefore(ctx *types.ServiceContext, project *setting.Project) error {
	log := p.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("project update before")
	return nil
}

func (p *Updater) UpdateAfter(ctx *types.ServiceContext, project *setting.Project) error {
	log := p.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("project update after")
	return nil
}
