package v1

import (
	"github.com/0x0034/tequila/api/v1/example"
)

type ApiGroup struct {
	ExampleGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
