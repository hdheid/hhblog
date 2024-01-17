package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/service/redis_ser"
)

func main() {
	core.InitConf()
	core.InitLogger()
	core.InitRedis()

	//redis_ser.Digg("6elR9osBUTARvWsVp8_C") //点赞操作
	fmt.Println(redis_ser.NewDigg().GetInfo()) //获取所有点赞数据

	//redis_ser.DiggClear()
}
