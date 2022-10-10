package template

import "bytes"

func FromCmdApi(projectName string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package cmd

import (
	_ "`)
	buffer.WriteString(projectName)
	buffer.WriteString(`/errcode"
	"`)
	buffer.WriteString(projectName)
	buffer.WriteString(`/route"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sunmi-OS/gocore/v2/conf/viper"
	"github.com/sunmi-OS/gocore/v2/utils"
	"github.com/sunmi-OS/gocore/v2/utils/closes"
	"github.com/urfave/cli/v2"
)

var Api = &cli.Command{
	Name:    "api",
	Aliases: []string{"a"},
	Usage:   "api start",
	Subcommands: []*cli.Command{
		{
			Name:   "start",
			Usage:  "开启运行api服务",
			Action: RunApi,
		},
	},
}

func RunApi(c *cli.Context) error {
	defer closes.Close()

	initConf()
	initDB()
	initCache()
	
	if utils.IsRelease() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 注册路由
	routes.Routes(r)

	err := endless.ListenAndServe(viper.C.GetString("network.ApiServiceHost")+":"+viper.C.GetString("network.ApiServicePort"), r)
	if err != nil {
		return err
	}
	return nil
}`)

}
