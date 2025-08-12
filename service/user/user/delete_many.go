package user

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type ManyDeleter struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *ManyDeleter) DeleteMany(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete many")
	return rsp, nil
}

func (u *ManyDeleter) DeleteManyBefore(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete many before")
	return nil
}

func (u *ManyDeleter) DeleteManyAfter(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete many after")
	return nil
}
