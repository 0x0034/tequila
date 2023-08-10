package initialize

import (
	"fmt"
	"github.com/robfig/cron/v3"

	"github.com/0x0034/tequila/config"
	"github.com/0x0034/tequila/global"
	"github.com/0x0034/tequila/utils"
)

func Timer() {
	if global.GlobalConfig.Timer.Start {
		for i := range global.GlobalConfig.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.GlobalConfig.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.GlobalTimer.AddTaskByFunc("ClearDB", global.GlobalConfig.Timer.Spec, func() {
					err := utils.ClearTable(global.GlobalDB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.GlobalConfig.Timer.Detail[i])
		}
	}
}
