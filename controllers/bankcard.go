package controllers

import (
	"github.com/astaxie/beego"
)

type BankcardController struct {
	beego.Controller
}

type Check4metaRespVo struct {
	Suc bool   `json:"suc"`
	Msg string `json:"msg"`
}

// @Summary 银行卡四要素检查
// @Description 银行卡四要素（身份证、姓名、手机号、银行卡号）检查
// @Success 200 {string} success
// @router /check4meta [post]
func (u *BankcardController) Check4meta() {
	/*msg := utils.HttpGet()
	isSuc := msg.Message.Result == "1"*/
	u.Data["json"] = Check4metaRespVo{true, ""}
	u.ServeJSON()
}
