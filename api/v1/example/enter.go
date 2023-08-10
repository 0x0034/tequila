package example

import "github.com/0x0034/tequila/service"

type ApiGroup struct {
	ExampleApi
}

var (
	runtimeService = service.ServiceGroupApp.ExampleServiceGroup.RuntimeService
)
