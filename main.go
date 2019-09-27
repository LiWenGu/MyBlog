package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"listblog/models"
	_ "listblog/routers"
	// 运行main方法之前 调用models包下的文件的init方法。先初始化数据库
	_ "listblog/models"
	"strings"
)

func main() {
	initTemplate()
	initSession()
	beego.Run()
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})
}

func initSession() {
	//beego的session序列号是用gob的方式，因此需要将注册models.User
	gob.Register(models.User{})
	//https://beego.me/docs/mvc/controller/session.md
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}
