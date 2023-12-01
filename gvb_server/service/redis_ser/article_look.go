package redis_ser

import (
	"gvb_server/global"
	"strconv"
)

const lookPrefix = "look"

// Look 浏览文章
func Look(id string) error {
	res, err := global.RDB.HGet(lookPrefix, id).Int()
	if err != nil {
		global.Log.Debug(err)
	}
	//如果redis没有该文章的hash，那么res是会等于0的，也可以直接进行下面的操作
	err = global.RDB.HSet(lookPrefix, id, res+1).Err()
	if err != nil {
		global.Log.Debug(err)
		return err
	}
	return nil
}

func GetLook(id string) (int, error) {
	res, err := global.RDB.HGet(lookPrefix, id).Int()
	if err != nil {
		return 0, err
	}

	return res, nil
}

// GetLookInfo 持久化，每隔一段时间都要去同步（存入es）
func GetLookInfo() map[string]int {
	var DiggInfo = make(map[string]int)

	diggs := global.RDB.HGetAll(lookPrefix).Val()
	for k, v := range diggs { //将string类型转换为数字类型
		num, _ := strconv.Atoi(v)
		DiggInfo[k] = num
	}

	return DiggInfo
}

func LookClear() {
	global.RDB.Del(lookPrefix) //直接将这个 hash 删除
}
