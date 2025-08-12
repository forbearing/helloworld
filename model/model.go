package model

import (
	"helloworld/model/setting"

	"github.com/forbearing/golib/model"
)

func Init() error {
	model.Register[*Group]()
	model.Register[*setting.Project]()
	model.Register[*User]()
	return nil
}
