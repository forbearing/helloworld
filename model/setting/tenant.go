package setting

import "github.com/forbearing/golib/model"

type Tenant struct {
	Name string `json:"name,omitempty" schema:"name"`

	model.Base
}
