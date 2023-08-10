package service

import (
	"github.com/0x0034/tequila/service/example"
)

type ServiceGroup struct {
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
