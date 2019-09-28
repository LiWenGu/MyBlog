package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/LiWenGu/payServer/controllers:BankcardController"] = append(beego.GlobalControllerRouter["github.com/LiWenGu/payServer/controllers:BankcardController"],
		beego.ControllerComments{
			Method:           "Check4meta",
			Router:           `/check4meta`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/LiWenGu/payServer/controllers:BankcardController"] = append(beego.GlobalControllerRouter["github.com/LiWenGu/payServer/controllers:BankcardController"],
		beego.ControllerComments{
			Method:           "FindOne",
			Router:           `/findOne`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
