package filter

import (
	"fmt"
	"github.com/LiWenGu/payServer/cache"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	var jwtFilter = func(ctx *context.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if len(token) == 0 || !CheckToken(token) {
			ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
			ctx.WriteString(`{"code":500}`)
		}
		cache.Put("liwenguang", "timeout")
		fmt.Print(cache.Get("liwenguang"))

	}
	beego.InsertFilter("/v1/bankcard/*)", beego.BeforeRouter, jwtFilter)

}
