package cache

import (
	"encoding/json"
	"fmt"
	"github.com/LiWenGu/payServer/controllers/vo"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"log"
	"time"
)

type RedisConfig struct {
	Key      string `json:"key"`
	Conn     string `json:"conn"`
	DbNum    string `json:"dbNum"`
	Password string `json:"password"`
}

var redis cache.Cache

func init() {
	config, err := json.Marshal(RedisConfig{
		beego.AppConfig.String("redisKey"),
		beego.AppConfig.String("redisConn"),
		beego.AppConfig.String("redisDbNum"),
		beego.AppConfig.String("redisPassword"),
	})
	redis, err = cache.NewCache("redis", string(config))
	if err != nil {
		log.Println(err)
	}
}

func Get(key string) string {
	return string(redis.Get(key).([]byte))
}

func PutStr(key string, value string) {
	d, _ := time.ParseDuration("1m")
	redis.Put(key, value, d)
}

func PutStruct(key *vo.TokenReqVo, value string) {
	b, err := json.Marshal(key)
	if err != nil {
		fmt.Println("error")
	}
	d, _ := time.ParseDuration("1m")
	redis.Put(string(b), value, d)
}
