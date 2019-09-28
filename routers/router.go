package routers

import (
	"github.com/astaxie/beego"
	"listblog/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
