package dns

import (
	"helloworld/model/cmdb"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Lister struct {
	service.Base[*cmdb.DNS, *cmdb.DNS, *cmdb.DNS]
}

func (d *Lister) List(ctx *types.ServiceContext, req *cmdb.DNS) (rsp *cmdb.DNS, err error) {
	log := d.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("dns list")
	return rsp, nil
}

func (d *Lister) ListBefore(ctx *types.ServiceContext, dns *[]*cmdb.DNS) error {
	log := d.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("dns list before")
	return nil
}

func (d *Lister) ListAfter(ctx *types.ServiceContext, dns *[]*cmdb.DNS) error {
	log := d.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("dns list after")
	return nil
}
