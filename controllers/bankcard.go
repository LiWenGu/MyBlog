package controllers

import (
	"encoding/json"
	"github.com/LiWenGu/payServer/controllers/vo"
	"github.com/LiWenGu/payServer/service"
	"github.com/LiWenGu/payServer/utils"
	"github.com/astaxie/beego"
)

// 银行卡相关操作
type BankcardController struct {
	beego.Controller
}

// @Summary 银行卡四要素检查
// @Description 银行卡四要素（身份证、姓名、手机号、银行卡号）检查
// @Param	body		body	vo.BankcardCheck4metaReqVo	true		"四要素"
// @Success 200 {object} vo.BankcardCheck4metaRespVo
// @router /check4meta [post]
func (u *BankcardController) Check4meta() {
	var req vo.BankcardCheck4metaReqVo
	json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	msg := utils.HttpGet(&req)
	isSuc := msg.Message.Result == "1"
	u.Data["json"] = vo.BankcardCheck4metaRespVo{isSuc, msg.Message.Code}
	u.ServeJSON()
}

// @Summary 查找银行卡
// @Description 银行卡四要素（身份证、姓名、手机号、银行卡号）检查
// @Success 200 {object} vo.BankcardCheck4metaRespVo
// @router /findOne [get]
func (u *BankcardController) FindOne() {
	bankService := service.BankService{}
	bankService.FindOne()
}
