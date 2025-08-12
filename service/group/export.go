package group

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Exporter struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *Exporter) Export(ctx *types.ServiceContext, data ...*model.Group) ([]byte, error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group export")
	return nil, nil
}
