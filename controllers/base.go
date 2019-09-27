package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"listblog/models"
	"log"
)

//定义session中的key值
const SESSION_USER_KEY = "SESSION_USER_KEY"

type NestPreparer interface {
	NestPreparer()
}

type BaseController struct {
	beego.Controller
	IsLogin bool        //标识 用户是否登陆
	User    models.User //登陆的用户
	Dao     *models.DB
}

func (ctx *BaseController) Prepare() {
	log.Println("BaseController")
	// 验证用户是否登陆，判断session中是否存在用户，存在就已经登陆，不存在就没有登陆。
	ctx.IsLogin = false
	tu := ctx.GetSession(SESSION_USER_KEY)
	if tu != nil {
		if u, ok := tu.(models.User); ok {
			ctx.User = u
			ctx.Data["User"] = u
			ctx.IsLogin = true
		}
	}
	ctx.Data["IsLogin"] = ctx.IsLogin
	// 将页面路径 保存到 Path变量里面
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	// 判断子类是否实现了NestPreparer接口，如果实现了就调用接口方法。
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPreparer()
	}
}

func (c *BaseController) GetMustString(key string, msg string) string {
	email := c.GetString(key, "")
	if len(email) == 0 {
		c.Abort500(errors.New(msg))
	}
	return email
}

func (ctx *BaseController) Abort500(err error) {
	ctx.Data["error"] = err
	ctx.Abort("500")
}

type H map[string]interface{}

type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (ctx *BaseController) JSONOk(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	ctx.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	ctx.ServeJSON()
}

func (ctx *BaseController) JSONOkH(msg string, maps H) {
	if maps == nil {
		maps = H{}
	}
	maps["code"] = 0
	maps["msg"] = msg
	ctx.Data["json"] = maps
	ctx.ServeJSON()
}

func (ctx *BaseController) JSONOkData(count int, data interface{}) {
	ctx.Data["json"] = &ResultJsonValue{
		Code:  0,
		Count: count,
		Msg:   "成功！",
		Data:  data,
	}
	ctx.ServeJSON()
}

func (this *BaseController) UUID() string {
	return "1"
}
