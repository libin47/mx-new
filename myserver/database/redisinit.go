package database

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdb redis.Client

func InitRDB(ctx context.Context) redis.Client {
	cfg := &redis.Options{
		Addr:     viper.GetString("redissource.host") + ":" + viper.GetString("redissource.port"),
		Password: viper.GetString("redissource.password"),
		DB:       viper.GetInt("redissource.DB"),
	}
	client := redis.NewClient(cfg)
	_, err := client.Ping(ctx).Result()

	if err != nil {
		fmt.Println("redis connect:", err)
	}
	rdb = *client
	return rdb
}

func GetRDB(ctx context.Context) redis.Client {
	return rdb
}

func GetRedisString(ctx context.Context, collection string, key string) string {
	key_all := collection + string('-') + key
	result, err := rdb.Get(ctx, key_all).Result()
	if err != nil {
		fmt.Println("redis get err:", err)
	}
	return result
}

func SetRedisString(ctx context.Context, collection string, key string, value string) bool {
	key_all := collection + string('-') + key
	err := rdb.Set(ctx, key_all, value, 0).Err()
	if err != nil {
		fmt.Println("redis set err:", err)
		return false
	}
	return true
}

func GetRedisInt(ctx context.Context, collection string, key string) int {
	key_all := collection + string('-') + key
	result, err := rdb.Get(ctx, key_all).Result()
	if err != nil {
		fmt.Println("redis get error:", err)
	}
	number, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println("redis get int error:", err)
		return 0
	}
	return number
}

func SetRedisInt(ctx context.Context, collection string, key string, value int) bool {
	key_all := collection + string('-') + key
	err := rdb.Set(ctx, key_all, value, 0).Err()
	if err != nil {
		fmt.Println("redis set error:", err)
		return false
	}
	return true
}
