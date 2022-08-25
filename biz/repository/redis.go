package repository

import (
	"github.com/go-redis/redis"
)

var RDB *redis.Client

//func init() {
//	RDB = redis.NewClient(&redis.Options{
//		Addr: utils.GetConfigResolve().GetConfigString("redis.address"),
//	})
//
//	_, err := RDB.Ping().Result()
//	if err != nil {
//		panic(err)
//	}
//}

func GetRedisDB() *redis.Client {
	return RDB
}
