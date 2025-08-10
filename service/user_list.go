package service

import (
	"helloworld/model"

	"github.com/forbearing/golib/types"
)

func (u *user) List(ctx *types.ServiceContext, req *model.User) (*model.User, error) {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user list")
	return &model.User{}, nil
}

func (u *user) ListBefore(ctx *types.ServiceContext, users *[]*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user list before")
	return nil
}

func (u *user) ListAfter(ctx *types.ServiceContext, users *[]*model.User) error {
	log := u.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("user list after")
	return nil
}
