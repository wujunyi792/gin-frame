package redis

import (
	"context"
	"github.com/wujunyi792/ginFrame/config"
	"github.com/wujunyi792/ginFrame/logger"
	"time"

	"github.com/go-redis/redis"
)

type wxRedis struct {
	Address string
	pClient *redis.Client
}

var sWxRedis *wxRedis
var gContext = context.Background()

func initFromConfig() {
	var address = config.Config["redis_address"]
	logger.Info.Printf("redis address:%s\n", address)
	var client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	sWxRedis = &wxRedis{Address: address, pClient: client}
	pong, err := client.Ping().Result()
	if err != nil {
		logger.Error.Println(err)
		return
	}
	logger.Info.Printf("redis ping result:%s\n", pong)
}

func GetWxRedis() *wxRedis {
	if sWxRedis == nil {
		initFromConfig()
	}
	return sWxRedis
}

func (r *wxRedis) Get(key string) (string, error) {
	if r.pClient == nil {
		logger.Error.Fatalln("pClient cannot be nil")
	}
	return r.pClient.Get(key).Result()
}

func (r *wxRedis) GetInt(key string) (int, error) {
	if r.pClient == nil {
		logger.Error.Fatalln("pClient cannot be nil")
	}
	//logger.Info.Printf("get val of key:%s", key)
	return r.pClient.Get(key).Int()
}

func (r *wxRedis) GetIntOrDefault(key string, def int) (int, error) {
	val, err := r.GetInt(key)
	if err == nil {
		return val, err
	}
	if err == redis.Nil {
		return def, nil //已处理redis.Nil异常
	}
	return def, err
}

func (r *wxRedis) Set(key string, value interface{}, expireDuration time.Duration) error {
	if r.pClient == nil {
		logger.Error.Fatalln("pClient is nil")
	}
	return r.pClient.Set(key, value, expireDuration).Err()
}

func (r *wxRedis) RemoveKey(key string, errWhenRedisNil bool) error {
	err := r.pClient.Del(key).Err()
	if err == redis.Nil && !errWhenRedisNil {
		return nil
	}
	return nil
}
