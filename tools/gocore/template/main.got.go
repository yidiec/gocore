package template

import "bytes"

func FromMain(projectName string, cmdList []string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package main

import (
	"os"

	"`)
	buffer.WriteString(projectName)
	buffer.WriteString(`/cmd"
	"`)
	buffer.WriteString(projectName)
	buffer.WriteString(`/conf"

	"github.com/sunmi-OS/gocore/v2/glog"
	"github.com/sunmi-OS/gocore/v2/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	// 打印Banner
	utils.PrintBanner(conf.ProjectName)
	// 配置cli参数
	app := cli.NewApp()
	app.Name = conf.ProjectName
	app.Usage = conf.ProjectName
	app.Version = conf.ProjectVersion

	// 指定命令运行的函数
	app.Commands = []*cli.Command{
        `)
	for _, cmd := range cmdList {
		buffer.WriteString(cmd + "\n")
	}
	buffer.WriteString(`
	}

	// 启动cli
	if err := app.Run(os.Args); err != nil {
		glog.ErrorF("Failed to start application: %v", err)
	}
}`)

}
