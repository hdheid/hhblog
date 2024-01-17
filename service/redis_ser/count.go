package redis_ser

import (
	"gvb_server/global"
	"strconv"
)

type CountDB struct {
	Index string //索引前缀
}

// Set 数据加一
func (c CountDB) Set(id string) error {
	res, err := global.RDB.HGet(c.Index, id).Int()
	if err != nil {
		global.Log.Debug(err)
	}
	//如果redis没有该文章的hash，那么res是会等于0的，也可以直接进行下面的操作
	err = global.RDB.HSet(c.Index, id, res+1).Err()
	if err != nil {
		global.Log.Debug(err)
		return err
	}
	return nil
}

// Get 取出数据
func (c CountDB) Get(id string) (int, error) {
	res, err := global.RDB.HGet(c.Index, id).Int()
	if err != nil {
		return 0, err
	}

	return res, nil
}

// GetInfo 取出数据信息
func (c CountDB) GetInfo() map[string]int {
	var DiggInfo = make(map[string]int)

	diggs := global.RDB.HGetAll(c.Index).Val()
	for k, v := range diggs { //将string类型转换为数字类型
		num, _ := strconv.Atoi(v)
		DiggInfo[k] = num
	}

	return DiggInfo
}

func (c CountDB) Clear() {
	global.RDB.Del(c.Index) //直接将这个 hash 删除
}
