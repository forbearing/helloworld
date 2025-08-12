package model

import (
	"helloworld/model/setting"

	"github.com/forbearing/golib/model"
)

func Init() error {
	model.Register[*cmdb.DNS]()
	model.Register[*cmdb.Machine]()
	model.Register[*Group]()
	model.Register[*setting.Project]()
	model.Register[*setting.Region]()
	model.Register[*setting.Tenant]()
	model.Register[*setting.Vendor]()
	model.Register[*User]()
	return nil
}
