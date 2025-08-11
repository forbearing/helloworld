package service

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type userUpdater struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *userUpdater) Update(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user update")
	return rsp, nil
}

func (u *userUpdater) UpdateBefore(ctx *types.ServiceContext, user *model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user update before")
	return nil
}

func (u *userUpdater) UpdateAfter(ctx *types.ServiceContext, user *model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user update after")
	return nil
}
