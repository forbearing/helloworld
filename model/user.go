package model

import (
	. "github.com/forbearing/golib/dsl"
	"github.com/forbearing/golib/model"
)

type User struct {
	Name string `json:"name"` // 用户名
	Age  int    `json:"age"`  // 用户年龄

	model.Base
}

type UserReq struct {
	Name string `json:"name"`
}

type UserRsp struct {
	Name string `json:"name"`
}

func (User) Design() {
	Enabled(true)
	// Endpoint("user")

	Create(func() {
		Enabled(true)
		// Payload[User]()
		// Result[User]()
		// Payload[UserReq]()
		// Result[UserRsp]()
	})
	Delete(func() {
		Enabled(true)
	})
	Update(func() {
		Enabled(true)
	})
	List(func() {
		Enabled(true)
	})
	Get(func() {
		Enabled(true)
	})
	CreateMany(func() {
		Enabled(true)
	})
	DeleteMany(func() {
		Enabled(true)
	})
	UpdateMany(func() {
		Enabled(true)
	})
	PatchMany(func() {
		Enabled(true)
	})
}
