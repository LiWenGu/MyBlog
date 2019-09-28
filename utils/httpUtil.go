package utils

import (
	"encoding/json"
	"github.com/LiWenGu/payServer/controllers/vo"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type BankcardCheck4meta struct {
	Code    int
	Value   string
	Message BankcardCheck4metaMsg
}

type BankcardCheck4metaMsg struct {
	Result string
	Code   string
	Tag    string
}

func HttpGet(metaReq *vo.BankcardCheck4metaReqVo) *BankcardCheck4meta {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://element4.market.alicloudapi.com/safrv_4meta?"+
		"__userId="+beego.AppConfig.String("userId")+"&"+
		"validator_id_num="+metaReq.Idnum+"&"+
		"validator_input_phone="+metaReq.Phone+"&"+
		"validator_user_bank_card="+metaReq.Bankcard+"&"+
		"validator_user_name="+metaReq.Name, nil)
	req.Header.Set("Authorization", "APPCODE "+beego.AppConfig.String("appCode"))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	result := BankcardCheck4meta{}
	json.Unmarshal(body, &result)
	return &result
}
