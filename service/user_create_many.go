package service

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type userManyCreator struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *userManyCreator) CreateMany(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user create many")
	return rsp, nil
}

func (u *userManyCreator) CreateManyBefore(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user create many before")
	return nil
}

func (u *userManyCreator) CreateManyAfter(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user create many after")
	return nil
}
