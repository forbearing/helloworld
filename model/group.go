package model

import "github.com/forbearing/golib/model"

type Group struct {
	Name  string
	Users []User

	model.Base
}
