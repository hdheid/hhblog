package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "wuwang222", //没有密码就设置为redis
		DB:       0,           //使用默认的DB
		PoolSize: 100,         //连接池大小
	})

	_, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel() //确保cancel()始终在函数返回的时候执行

	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Error(err)
		return
	}
}

func main() {
	err := rdb.Set("name", "glz", 10*time.Second).Err()
	if err != nil {
		logrus.Error(err)
		return
	}
	cmd := rdb.Keys("name")
	keys, err := cmd.Result()
	println(keys)
}
