package machine

import (
	"helloworld/model/cmdb"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Lister struct {
	service.Base[*cmdb.Machine, *cmdb.Machine, *cmdb.Machine]
}

func (m *Lister) List(ctx *types.ServiceContext, req *cmdb.Machine) (rsp *cmdb.Machine, err error) {
	log := m.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("machine list")
	return rsp, nil
}

func (m *Lister) ListBefore(ctx *types.ServiceContext, machines *[]*cmdb.Machine) error {
	log := m.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("machine list before")
	return nil
}

func (m *Lister) ListAfter(ctx *types.ServiceContext, machines *[]*cmdb.Machine) error {
	log := m.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("machine list after")
	return nil
}
