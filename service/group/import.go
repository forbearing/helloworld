package group

import (
	"helloworld/model"
	"io"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Importer struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *Importer) Import(ctx *types.ServiceContext, reader io.Reader) ([]*model.Group, error) {
	log := g.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("group import")
	return nil, nil
}
