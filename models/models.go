package models

import "github.com/astaxie/beego/orm"

import (
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:Zhuangdian!@#@tcp(118.31.51.165:3306)/mall?charset=utf8")
}
