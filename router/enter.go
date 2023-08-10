package router

import "github.com/0x0034/tequila/router/example"

type RouterGroup struct {
	Example example.ExampleRouter
}

var RouterGroupApp = new(RouterGroup)
