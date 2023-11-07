package core

import (
	"context"
	"github.com/go-redis/redis"
	"gvb_server/global"
	"time"
)

func InitRedis() {
	RedisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisConf.Addr(),
		Password: RedisConf.Password, //没有密码就设置为redis
		DB:       0,                  //使用默认的DB
		PoolSize: RedisConf.PoolSize, //连接池大小
	})

	_, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel() //确保cancel()始终在函数返回的时候执行

	_, err := rdb.Ping().Result()
	if err != nil {
		global.Log.Errorf("Redis 连接失败 %s", err)
		return
	}

	global.Log.Info("Redis 连接成功!")
	global.RDB = rdb
}
