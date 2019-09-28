package service

import (
	"github.com/LiWenGu/payServer/models"
)

// 银行卡相关操作
type BankService struct {
}

func (*BankService) FindOne() {
	models.FindOne()
}
