package service

import (
	"github.com/forbearing/golib/service"
)

func Init() error {
	service.Register[*user]()
	return nil
}
