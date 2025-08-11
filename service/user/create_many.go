package user

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type ManyCreator struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *ManyCreator) CreateMany(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user create many")
	return rsp, nil
}

func (u *ManyCreator) CreateManyBefore(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user create many before")
	return nil
}

func (u *ManyCreator) CreateManyAfter(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user create many after")
	return nil
}
