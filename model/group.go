package model

import (
	. "github.com/forbearing/golib/dsl"
	"github.com/forbearing/golib/model"
	"go.uber.org/zap/zapcore"
)

type Group struct {
	Name        string `json:"name,omitempty" schema:"name"`  // 用户组名
	Desc        string `json:"desc,omitempty" schema:"desc"`  // 用户组描述
	MemberCount int    `json:"member_count" gorm:"default:0"` // 成员数量

	model.Base
}

func (g *Group) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if g == nil {
		return nil
	}
	enc.AddString("name", g.Name)
	enc.AddString("desc", g.Desc)
	enc.AddInt("member_count", g.MemberCount)
	enc.AddObject("base", &g.Base)
	return nil
}

func (Group) Design() {
	// Enabled(false)

	// Create(func() {
	// 	Enabled(true)
	// })
	// Delete(func() {
	// 	Enabled(true)
	// })
	// Update(func() {
	// 	Enabled(true)
	// })
	// Patch(func() {
	// 	Enabled(true)
	// })
	// List(func() {
	// 	Enabled(true)
	// })
	Get(func() {
		Enabled(true)
	})

	CreateMany(func() {
		Enabled(true)
	})

	// DeleteMany(func() {
	// 	Enabled(true)
	// })
	// UpdateMany(func() {
	// 	Enabled(true)
	// })
	// PatchMany(func() {
	// 	Enabled(true)
	// })
}
