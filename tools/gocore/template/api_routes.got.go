// Code generated by hero.
// source: /Users/SM0286/code/core/gocore/tools/gocore/template/api_routes.got
// DO NOT EDIT!
package template

import (
	"bytes"

	"github.com/shiyanhui/hero"
)

func FromApiRoutes(pkg, routes string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package routes

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	`)
	hero.EscapeHTML(pkg, buffer)
	buffer.WriteString(`
	"github.com/labstack/echo/v4"
	"github.com/sunmi-OS/gocore/v2/utils/log"
)

var (
	pid      int
	progname string
)

func init() {
	pid = os.Getpid()
	paths := strings.Split(os.Args[0], "/")
	paths = strings.Split(paths[len(paths)-1], string(os.PathSeparator))
	progname = paths[len(paths)-1]

}

// Router 初始化路由
func Router(e *echo.Echo) {

	// 内存溢出检测
	e.Any("/pprof-init", func(context echo.Context) error {
		pid = os.Getpid()
		paths := strings.Split(os.Args[0], "/")
		paths = strings.Split(paths[len(paths)-1], string(os.PathSeparator))
		progname = paths[len(paths)-1]
		runtime.MemProfileRate = 1
		return nil
	})
	// 内存溢出检测
	e.Any("/pprof", func(context echo.Context) error {
		runtime.GC()
		f, err := os.Create(fmt.Sprintf("./heap_%s_%d_%s.prof", progname, pid, time.Now().Format("2006_01_02_03_04_05")))
		if err != nil {
			return err
		}
		defer f.Close()
		err = pprof.Lookup("heap").WriteTo(f, 1)
		if err != nil {
			log.Sugar.Info(err)
		}
		runtime.MemProfileRate = 0
		return context.JSON(200, "pong")
	})

    `)
	hero.EscapeHTML(routes, buffer)
	buffer.WriteString(`
}
`)

}