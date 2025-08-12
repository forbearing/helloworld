package user

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type ManyPatcher struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *ManyPatcher) PatchMany(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user patch many")
	return rsp, nil
}

func (u *ManyPatcher) PatchManyBefore(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user patch many before")
	return nil
}

func (u *ManyPatcher) PatchManyAfter(ctx *types.ServiceContext, users ...*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user patch many after")
	return nil
}
