package example

import (
	v1 "github.com/0x0034/tequila/api/v1"
	"github.com/gin-gonic/gin"
)

type ExampleRouter struct{}

func (e *ExampleRouter) InitRuntimeRouter(Router *gin.RouterGroup) {
	exampleRouter := Router.Group("example")
	exampleService := v1.ApiGroupApp.ExampleGroup.ExampleApi
	{
		exampleRouter.GET("getRuntime", exampleService.GetRuntime)

	}

}
