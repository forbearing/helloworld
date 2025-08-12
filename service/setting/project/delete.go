package project

import (
	"helloworld/model/setting"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Deleter struct {
	service.Base[*setting.Project, *setting.Project, *setting.Project]
}

func (p *Deleter) Delete(ctx *types.ServiceContext, req *setting.Project) (rsp *setting.Project, err error) {
	log := p.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("project delete")
	return rsp, nil
}

func (p *Deleter) DeleteBefore(ctx *types.ServiceContext, project *setting.Project) error {
	log := p.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("project delete before")
	return nil
}

func (p *Deleter) DeleteAfter(ctx *types.ServiceContext, project *setting.Project) error {
	log := p.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("project delete after")
	return nil
}
