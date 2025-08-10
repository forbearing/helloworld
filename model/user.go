package model

import (
	. "github.com/forbearing/golib/dsl"
	"github.com/forbearing/golib/model"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	model.Base
}

func (User) Design() {
	Create(func() {
		Enabled(true)
		Payload[User]()
		Result[User]()
	})

	List(func() {
		Enabled(true)
	})
}
