package template

import "bytes"

func FromCmdCronJob(name, cronjobs string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package cmd

import (
	"`)
	buffer.WriteString(name)
	buffer.WriteString(`/cronjob"
	"github.com/robfig/cron/v3"
	"github.com/sunmi-OS/gocore/v2/utils/closes"
	"github.com/urfave/cli/v2"
)

var Cron = &cli.Command{
	Name:    "cron",
	Aliases: []string{"c"},
	Usage:   "cron start",
	Subcommands: []*cli.Command{
		{
			Name:   "start",
			Usage:  "开启运行api服务",
			Action: CronTable,
		},
	},
}

func CronTable(c *cli.Context) error {
	defer closes.Close()
	// 初始化必要内容
	initConf()
	initDB()
	initCache()

	cronJob := cron.New()

    `)
	buffer.WriteString(cronjobs)
	buffer.WriteString(`

	cronJob.Start()

	closes.AddShutdown(closes.ModuleClose{
		Name:     "CronTable",
		Priority: 0,
		Func: func() {
			cronJob.Stop()
		},
	})
	closes.SignalClose()
	return nil
}`)

}
