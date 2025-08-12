package region

import (
	"helloworld/model/setting"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Lister struct {
	service.Base[*setting.Region, *setting.Region, *setting.Region]
}

func (r *Lister) List(ctx *types.ServiceContext, req *setting.Region) (rsp *setting.Region, err error) {
	log := r.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("region list")
	return rsp, nil
}

func (r *Lister) ListBefore(ctx *types.ServiceContext, regions *[]*setting.Region) error {
	log := r.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("region list before")
	return nil
}

func (r *Lister) ListAfter(ctx *types.ServiceContext, regions *[]*setting.Region) error {
	log := r.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("region list after")
	return nil
}
