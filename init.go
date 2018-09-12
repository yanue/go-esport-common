/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : init.go
 Time    : 2018/9/12 19:18
 Author  : yanue
 
 - 初始化操作:
 - 初始化env参数
 
------------------------------- go ---------------------------------*/

package common

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
)

func init() {
	// 增加启动参数env,参数值:dev|test|prod
	app := cmd.App()

	app.Flags = append(app.Flags, cli.StringFlag{
		Name:  "env",
		Usage: "run environment (dev or test or prod)",
	})

	before := app.Before

	app.Before = func(ctx *cli.Context) error {
		if env := ctx.String("env"); len(env) > 0 {
			// got config
			// do stuff
			ConfigEnv = env
		}

		// 获取env参数后
		// 重新初始化日志
		initLogs()

		return before(ctx)
	}
}
