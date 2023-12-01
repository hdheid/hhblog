package redis_ser

import (
	"gvb_server/global"
	"strconv"
)

const diggPrefix = "digg"

// Digg 点赞操作
func Digg(id string) error {
	res, err := global.RDB.HGet(diggPrefix, id).Int()
	if err != nil {
		global.Log.Debug(err)
	}
	//如果redis没有该文章的hash，那么res是会等于0的，也可以直接进行下面的操作
	err = global.RDB.HSet(diggPrefix, id, res+1).Err()
	if err != nil {
		global.Log.Debug(err)
		return err
	}
	return nil
}

func GetDigg(id string) (int, error) {
	res, err := global.RDB.HGet(diggPrefix, id).Int()
	if err != nil {
		return 0, err
	}

	return res, nil
}

// GetDiggsInfo 持久化，每隔一段时间都要去同步（存入es）
func GetDiggsInfo() map[string]int {
	var DiggInfo = make(map[string]int)

	diggs := global.RDB.HGetAll(diggPrefix).Val()
	for k, v := range diggs { //将string类型转换为数字类型
		num, _ := strconv.Atoi(v)
		DiggInfo[k] = num
	}

	return DiggInfo
}

func DiggClear() {
	global.RDB.Del(diggPrefix) //直接将这个 hash 删除
}
