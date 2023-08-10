package example

import (
	"github.com/0x0034/tequila/models/common/response"
	"github.com/gin-gonic/gin"
)

type ExampleApi struct{}

// GetRuntime
// @Tags      RuntimeApi
// @Summary   获取当前运行模式
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=string,msg=string}  "创建基础api"
// @Router    /baileys/example/getRuntime [get]
func (e ExampleApi) GetRuntime(c *gin.Context) {
	podList := runtimeService.GetRuntime()
	response.OkWithData(podList, c)
}
