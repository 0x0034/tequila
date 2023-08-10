package example

import "github.com/0x0034/tequila/global"

type RuntimeService struct{}

func (r RuntimeService) GetRuntime() (runtime string) {
	return global.GlobalConfig.Mysql.LogMode
}
