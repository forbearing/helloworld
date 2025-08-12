package setting

import (
	"github.com/forbearing/golib/dsl"
	"github.com/forbearing/golib/model"
)

type Region struct {
	Name string `json:"name,omitempty" schema:"name"`

	model.Base
}

func (Region) Design() {
	dsl.Create(func() {
		dsl.Enabled(true)
	})
	dsl.List(func() {
		dsl.Enabled(true)
	})
}
