package service

import (
	"helloworld/service/group"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types/consts"
)

func Init() error {
	service.Register[*group.Getter](consts.PHASE_GET)
	service.Register[*group.ManyCreator](consts.PHASE_CREATE_MANY)
	return nil
}
