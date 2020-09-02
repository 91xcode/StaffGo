package redis

import (
	"gopkg.in/redis.v5"
)

var (
	ErrNil = redis.Nil
	client *redis.Client
)

func Init_Test(opt *redis.FailoverOptions) {
	client = redis.NewFailoverClient(opt)
}

func Init(opt *redis.Options) {
	client = redis.NewClient(opt)
}

func Close() {
	client.Close()
}

func Ping() error {
	return client.Ping().Err()
}
