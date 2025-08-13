package twofa

import (
	. "github.com/forbearing/golib/dsl"
	"github.com/forbearing/golib/model"
)

type Setup struct {
	model.Base
}

type SetupRsp = []byte

func (Setup) Design() {
	Migrate(false)
	List(func() {
		Enabled(true)
		Result[SetupRsp]()
	})
}
