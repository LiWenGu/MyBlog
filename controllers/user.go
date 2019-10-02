package controllers

import (
	"github.com/LiWenGu/payServer/controllers/vo"
	"github.com/LiWenGu/payServer/filter"
	"github.com/astaxie/beego"
)

// 银行卡相关操作
type UserController struct {
	beego.Controller
}

// @Summary 获取Token
// @Description 获取Token
// @Success 200 {object} vo.Token
// @router /getToken [get]
func (u *UserController) GetToken() {
	u.Data["json"] = vo.Token{200, filter.GenToken()}
	u.ServeJSON()
}
