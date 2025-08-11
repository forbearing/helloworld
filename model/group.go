package model

import (
	. "github.com/forbearing/golib/dsl"
	"github.com/forbearing/golib/model"
)

type Group struct {
	Name  string
	Users []User

	model.Base
}

func (Group) Design() {
	// Enabled(false)

	Create(func() {
		Enabled(true)
	})
	Delete(func() {
		Enabled(true)
	})
	Update(func() {
		Enabled(true)
	})
	Patch(func() {
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
