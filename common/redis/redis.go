package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()
var rdb *redis.Client

func GetRdb() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.addr"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		})
	}
	return rdb
}

// 通用
func Exists(key string) bool {
	i, err := GetRdb().Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return i > 0
}

func TTL(key string) time.Duration {
	t, err := GetRdb().TTL(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return t
}

func Expire(key string, expire time.Duration) {
	err := GetRdb().Expire(ctx, key, expire).Err()
	if err != nil {
		panic(err)
	}
}

func Del(key string) {
	err := GetRdb().Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
}

// 字符串
func Set(key string, value string) {
	err := GetRdb().Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func SetEX(key string, value string, expiration time.Duration) {
	err := GetRdb().SetEX(ctx, key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	value, err := GetRdb().Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return value
}

// 哈希
func HSet(key string, field string, value interface{}) {
	err := GetRdb().HSet(ctx, key, field, value).Err()
	if err != nil {
		panic(err)
	}
}

func HGet(key string, field string) string {
	value, err := GetRdb().HGet(ctx, key, field).Result()
	if err != nil {
		panic(err)
	}
	return value
}

func HGetAll(key string) map[string]string {
	stringMap, err := GetRdb().HGetAll(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return stringMap
}
