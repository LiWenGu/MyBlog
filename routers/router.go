package routers

import (
	"github.com/astaxie/beego"
	"listblog/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(
		&controllers.IndexController{},
		&controllers.UserController{},
	)
}
