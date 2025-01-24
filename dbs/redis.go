package dbs

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Red *redis.Client

func InitRedis() {
	Red = redis.NewClient(
		&redis.Options{
			Addr:         viper.GetString("redis.addr"),
			Password:     viper.GetString("redis.password"),
			DB:           viper.GetInt("redis.DB"),
			PoolSize:     viper.GetInt("redis.poolSize"),
			MinIdleConns: viper.GetInt("redis.minIdleConn"),
		})
}
