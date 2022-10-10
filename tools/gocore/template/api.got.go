package template

import "bytes"

func FromApi(name, handler, apiContent string, comments []string, functions []string, req []string, buffer *bytes.Buffer) {
	if apiContent == "" {
		buffer.WriteString(`
package api

import (
	"`)
		buffer.WriteString(name)
		buffer.WriteString(`/param"

	"github.com/gin-gonic/gin"
    "github.com/sunmi-OS/gocore/v2/api"
)
`)
	} else {
		buffer.WriteString(apiContent)
	}
	for k1, v1 := range functions {
		buffer.WriteString(`
	// `)
		buffer.WriteString(v1)
		buffer.WriteString(" " + comments[k1])
		buffer.WriteString(`
    func `)
		buffer.WriteString(v1)
		buffer.WriteString(`(g *gin.Context) {
        ctx := api.NewContext(g)
        req := new(param.`)
		buffer.WriteString(req[k1])
		buffer.WriteString(`Request)
        err := ctx.BindValidator(req)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.Success(param.`)
		buffer.WriteString(req[k1])
		buffer.WriteString(`Response{})
    }
`)
	}

}
