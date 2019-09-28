package vo

type BankcardCheck4metaRespVo struct {
	Suc bool   `json:"suc" description:"是否成功"`
	Msg string `json:"msg" description:"如果失败，会返回错误code"`
}
