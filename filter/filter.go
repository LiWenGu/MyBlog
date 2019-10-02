package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func init() {
	var jwtFilter = func(ctx *context.Context) {
		token := ctx.Request.Header.Get("Authorization")
		logs.Info(token + "请求访问")
		if len(token) == 0 || !CheckToken(token) {
			ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
			ctx.WriteString(`{"code":500}`)
		}
	}
	beego.InsertFilter("/v1/bankcard/*)", beego.BeforeRouter, jwtFilter)
}
