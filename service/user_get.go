package service

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type userGetter struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *userGetter) Get(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user get")
	return rsp, nil
}

func (u *userGetter) GetBefore(ctx *types.ServiceContext, user *model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user get before")
	return nil
}

func (u *userGetter) GetAfter(ctx *types.ServiceContext, user *model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user get after")
	return nil
}
