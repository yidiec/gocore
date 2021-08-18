// Code generated by hero.
// source: /Users/SM0286/code/core/gocore/tools/gocore/template/model.got
// DO NOT EDIT!
package template

import (
	"bytes"
	"strings"

	"github.com/shiyanhui/hero"
)

func FromModel(dbName, tabels string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package `)
	hero.EscapeHTML(dbName, buffer)
	buffer.WriteString(`

import (
	"fmt"
"` + goCoreConfig.Service.ProjectName + `/conf"
	"gorm.io/gorm"
	g "github.com/sunmi-OS/gocore/v2/db/orm"
	"github.com/sunmi-OS/gocore/v2/utils"
)

func Orm() *gorm.DB {
	db := g.GetORM(conf.DB`)
	hero.EscapeHTML(strings.Title(dbName), buffer)
	buffer.WriteString(`)
	if utils.GetRunTime() != "onl" {
		db = db.Debug()
	}
	return db
}

func SchemaMigrate() {
	fmt.Println("开始初始化` + dbName + `数据库")
	//自动建表，数据迁移
    `)
	buffer.WriteString(tabels)
	buffer.WriteString(`
	fmt.Println("数据库` + "`" + ` + dbName + ` + "`" + `初始化完成")
}`)

}
