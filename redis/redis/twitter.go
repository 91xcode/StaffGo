package redis

import (
	"encoding/json"
	"code.be.staff.com/staff/StaffGo/public/zap"
	twittermodel "code.be.staff.com/staff/StaffGo/redis/model/twitter"
	"time"
)

type TwitterRedisDao struct {
	tokenKeyPrefix string
}

// NewUserRedisDao new a UserRedisDao and return.
func NewTwitterRedisDao() *TwitterRedisDao {
	return &TwitterRedisDao{
		tokenKeyPrefix: "twitter_",
	}
}

func (dao *TwitterRedisDao) SetToken(requestToken *twittermodel.TokenInfo, expire time.Duration) (err error) {
	bs, err := json.Marshal(requestToken)
	if err != nil {
		zap.Zlogger.Infof("json.Marshal(%v) error(%v)", requestToken, err)
		return
	}

	if err := client.Set(dao.tokenKeyPrefix+requestToken.Token, bs, expire).Err(); err != nil {
		zap.Zlogger.Infof("client.Set(%v,%v,%v) error(%v)", dao.tokenKeyPrefix+requestToken.Token, expire, string(bs), err)
	}
	zap.Zlogger.Infof("client.Set(%v,%v,%v)", dao.tokenKeyPrefix+requestToken.Token, expire, string(bs))
	return
}

func (dao *TwitterRedisDao) GetToken(token string)(requestToken *twittermodel.TokenInfo, err error)  {
	key:=dao.tokenKeyPrefix+token
	var bs []byte
	bs, err = client.Get(key).Bytes()
	if err != nil {
		if err == ErrNil {
			return nil, nil
		}
		zap.Zlogger.Infof("client.Get(%v) error(%v)", key, err)
		return nil, err
	}
	zap.Zlogger.Infof(string(bs))
	requestToken = &twittermodel.TokenInfo{}
	err = json.Unmarshal(bs, requestToken)
	if err != nil {
		zap.Zlogger.Infof("json.Unmarshal(%s) error(%v)", string(bs), err)
		client.Del(key)
	}
	return requestToken,err
}

func (dao *TwitterRedisDao) RemoveToken(requestToken *twittermodel.TokenInfo,)(err error)  {
	key:=dao.tokenKeyPrefix+requestToken.Token
	if err := client.Del(key).Err(); err != nil {
		zap.Zlogger.Infof("client.Del(%v) error(%v)", key, err)
		return err
	}
	return nil
}
