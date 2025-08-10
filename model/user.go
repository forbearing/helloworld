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

type UserReq struct {
	Name string `json:"name"`
}

type UserRsp struct {
	Name string `json:"name"`
}

func (User) Design() {
	Create(func() {
		Enabled(true)
		Payload[User]()
		Result[User]()
		// Payload[UserReq]()
		// Result[UserRsp]()
	})

	List(func() {
		Enabled(false)
	})
}
