package service

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type userDeleter struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *userDeleter) Delete(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete")
	return rsp, nil
}

func (u *userDeleter) DeleteBefore(ctx *types.ServiceContext, user *model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete before")
	return nil
}

func (u *userDeleter) DeleteAfter(ctx *types.ServiceContext, user *model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete after")
	return nil
}
