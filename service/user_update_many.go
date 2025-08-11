package service

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type userManyUpdater struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *userManyUpdater) UpdateMany(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user update many")
	return rsp, nil
}

func (u *userManyUpdater) UpdateManyBefore(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user update many before")
	return nil
}

func (u *userManyUpdater) UpdateManyAfter(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user update many after")
	return nil
}
