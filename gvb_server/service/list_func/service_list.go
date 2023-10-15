package list_func

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool //是否展示sql语句
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	//判断是否想展示所有sql语句
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLogger})
	}

	//设置默认排序，按照时间从晚到早排序
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按照时间往前排序
	}

	//将图片全部取出
	count = DB.Select("id").Find(&list).RowsAffected //count为总数
	global.Log.Debug(count)

	offset := (option.Page - 1) * option.Limit //分页查询的某一个公式
	if offset < 0 {                            //防止Page为0的时候
		offset = 0
	}

	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error //分页查询,加上sort排序
	if err != nil {
		global.Log.Debug("分页查询数据库处理错误，", err)
	}

	return list, count, err
}
