package main

import (
	"os"
	"code.be.staff.com/staff/StaffGo/public/zap"
	"code.be.staff.com/staff/StaffGo/redis/conf"
	"time"

	twittermodel "code.be.staff.com/staff/StaffGo/redis/model/twitter"
	"code.be.staff.com/staff/StaffGo/redis/redis"
)

func Init() {
	zap.Init(&conf.Conf.Zap)
	zap.Zlogger.Info("zap init success")

	redis.Init(&conf.Conf.Redis)
	zap.Zlogger.Info("redis init success")
}

func main() {
	if len(os.Args) < 2 {
		panic("params error you need  go run main.go config.ini")
	}

	if err := conf.Init(os.Args[1]); err != nil {
		panic(err)
	}

	Init()

	twitterRedisDao := redis.NewTwitterRedisDao()

	//set
	ttl := time.Minute * 10
	zap.Zlogger.Infof("ttl:%v", ttl)

	data := twittermodel.TokenInfo{
		Token:  "65692284671",
		Secret: "aktest",
	}
	rst := twitterRedisDao.SetToken(&data, ttl)
	zap.Zlogger.Infof("rst:%v", rst)

	//get
	token := "65692284671"
	result, err := twitterRedisDao.GetToken(token)
	if err != nil {
		zap.Zlogger.Infof("twitterRedisDao.GetToken(%v) error(%v)", token, err)
	}
	zap.Zlogger.Infof("result:%v", result)

	//remove
	ret:=twitterRedisDao.RemoveToken(&data)
	zap.Zlogger.Infof("ret:%v", ret)
}
