package group

import (
	"fmt"
	"helloworld/model"

	"github.com/forbearing/golib/database"
	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
)

type Creator struct {
	service.Base[*model.Group, *model.Group, *model.Group]
}

func (g *Creator) Create(ctx *types.ServiceContext, req *model.Group) (rsp *model.Group, err error) {
	g.Logger.Info("group create")
	if err := database.Database[*model.Group]().Create(&model.Group{Name: req.Name}); err != nil {
		return rsp, err
	}
	rsp = &model.GroupRsp{NameCustom: req.Name, DescCustom: "haha", Count: 1}
	return rsp, nil
}

func (g *Creator) CreateBefore(ctx *types.ServiceContext, group *model.Group) error {
	return nil
}

func (g *Creator) CreateAfter(ctx *types.ServiceContext, group *model.Group) error {
	fmt.Println("g.Logger is nil", g.Logger == nil)
	return nil
}
