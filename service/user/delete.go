package user

import (
	"helloworld/model"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Deleter struct {
	service.Base[*model.User, *model.User, *model.User]
}

func (u *Deleter) Delete(ctx *types.ServiceContext, req *model.User) (rsp *model.User, err error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete")
	return rsp, nil
}

func (u *Deleter) DeleteBefore(ctx *types.ServiceContext, user *model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete before")
	return nil
}

func (u *Deleter) DeleteAfter(ctx *types.ServiceContext, user *model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user delete after")
	return nil
}
