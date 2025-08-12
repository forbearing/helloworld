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

type GroupReq struct {
	Name string `json:"name,omitempty" schema:"name"`
}
type GroupRsp struct {
	NameCustom string   `json:"custom_name,omitempty" schema:"custom_name"`
	DescCustom string   `json:"custom_desc,omitempty" schema:"custom_desc"`
	Count      int      `json:"custom_count,omitempty" schema:"custom_count"`
	Deleted    []*Group `json:"deleted,omitempty"`
}

func (Group) Design() {
	// Enabled(false)

	Create(func() {
		Enabled(true)
		Payload[*GroupReq]()
		Result[GroupRsp]()
	})
	Delete(func() {
		Enabled(true)
		Payload[*GroupReq]()
		Result[*GroupRsp]()
	})
	Update(func() {
		Enabled(true)
		Payload[*GroupReq]()
		Result[*GroupRsp]()
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
	Import(func() {
		Enabled(true)
	})

	Export(func() {
		Enabled(true)
	})
}
