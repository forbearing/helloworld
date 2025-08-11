package service

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type userLister struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *userLister) List(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user list")
	return rsp, nil
}

func (u *userLister) ListBefore(ctx *types.ServiceContext, users *[]*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user list before")
	return nil
}

func (u *userLister) ListAfter(ctx *types.ServiceContext, users *[]*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user list after")
	return nil
}
