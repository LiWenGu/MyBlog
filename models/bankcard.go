package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 银行卡
type Bankcard struct {
	BankcardId int    `orm:"PK;column(BankcardId)"`
	CardNo     string `orm:"column(CardNo)"`
}

func (u *Bankcard) TableName() string {
	return "Bankcard"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Bankcard))
}

func FindOne() {
	o := orm.NewOrm()
	bankcard := Bankcard{BankcardId: 12}

	err := o.Read(&bankcard)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(bankcard.BankcardId, bankcard.CardNo)
	}
}
