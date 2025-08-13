package twofa

import (
	. "github.com/forbearing/golib/dsl"
	"github.com/forbearing/golib/model"
)

type Verify struct {
	model.Base
}

type VerifyReq struct {
	User string `json:"user"`
	Code string `json:"code"`
}

func (Verify) Design() {
	Migrate(false)
	Create(func() {
		Enabled(true)
		Payload[VerifyReq]()
	})
}
