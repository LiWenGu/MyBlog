package controllers

import (
	"encoding/json"
	"github.com/LiWenGu/payServer/cache"
	"github.com/LiWenGu/payServer/controllers/vo"
	"github.com/LiWenGu/payServer/filter"
	"github.com/astaxie/beego"
)

// 银行卡相关操作
type TokenController struct {
	beego.Controller
}

// @Summary 获取Token
// @Description 获取Token
// @Success 200 {object} vo.Token
// @router /getToken [get]
func (u *TokenController) GetToken() {
	var req vo.TokenReqVo
	json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	token := filter.GenToken()
	u.Data["json"] = vo.TokenRespVo{200, token}
	cache.PutStruct(&req, token)
	u.ServeJSON()
}
