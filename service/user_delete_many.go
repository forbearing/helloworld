package service

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type userManyDeleter struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *userManyDeleter) DeleteMany(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete many")
	return rsp, nil
}

func (u *userManyDeleter) DeleteManyBefore(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete many before")
	return nil
}

func (u *userManyDeleter) DeleteManyAfter(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete many after")
	return nil
}
