package setting

import (
	. "github.com/forbearing/golib/dsl"
	"github.com/forbearing/golib/model"
)

type Project struct {
	Name string `json:"name"`

	model.Base
}

func (Project) Design() {
	Enabled(true)
	Create(func() {
		Enabled(true)
	})
	Update(func() {
		Enabled(true)
	})
	Delete(func() {
		Enabled(true)
	})
}
