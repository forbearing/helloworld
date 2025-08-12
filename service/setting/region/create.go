package region

import (
	"helloworld/model/setting"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Creator struct {
	service.Base[*setting.Region, *setting.Region, *setting.Region]
}

func (r *Creator) Create(ctx *types.ServiceContext, req *setting.Region) (rsp *setting.Region, err error) {
	log := r.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("region create")
	return rsp, nil
}

func (r *Creator) CreateBefore(ctx *types.ServiceContext, region *setting.Region) error {
	log := r.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("region create before")
	return nil
}

func (r *Creator) CreateAfter(ctx *types.ServiceContext, region *setting.Region) error {
	log := r.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("region create after")
	return nil
}
