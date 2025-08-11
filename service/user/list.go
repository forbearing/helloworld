package user

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Lister struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *Lister) List(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user list")
	return rsp, nil
}

func (u *Lister) ListBefore(ctx *types.ServiceContext, users *[]*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user list before")
	return nil
}

func (u *Lister) ListAfter(ctx *types.ServiceContext, users *[]*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user list after")
	return nil
}
