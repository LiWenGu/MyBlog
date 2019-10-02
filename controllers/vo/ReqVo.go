package vo

type BankcardCheck4metaReqVo struct {
	Idnum    string `description:"身份证号"`
	Phone    string `description:"预留手机号"`
	Bankcard string `description:"银行卡账号"`
	Name     string `description:"持卡人姓名"`
}

type Token struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}