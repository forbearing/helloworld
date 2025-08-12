package model

import "github.com/forbearing/golib/model"

func Init() error {
	model.Register[*Group]()
	model.Register[*User]()
	return nil
}
