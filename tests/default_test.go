package test

import (
	_ "PayServer/routers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://element4.market.alicloudapi.com/safrv_4meta?"+
		"__userId=1961444063236731&"+
		"validator_id_num=429004199601072193&"+
		"validator_input_phone=17671656706&"+
		"validator_user_bank_card=6214832709883503&"+
		"validator_user_name=李文广", nil)
	req.Header.Set("Authorization", "APPCODE 78dc692db1114d358ffae4ad4f5ef8a3")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}

func TestParse(t *testing.T) {
	a := `{"code":200,"value":"1","message":{"result":"1","tag":"Y"}}`
	b := BankcardCheck4meta{}
	json.Unmarshal([]byte(a), &b)
	fmt.Println(b.Message)
}

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
