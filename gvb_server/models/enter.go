package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` //主键ID
	CreatedAt time.Time `json:"created_at"`           //创建时间
	UpdatedAt time.Time `json:"-"`                    //更新时间
}

// RemoveRequest 删除列表
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ESIdRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIdListRequest struct {
	IDList []string `json:"id_list" form:"id" uri:"id"`
}

type PageInfo struct {
	Page  int    `form:"page"`  //标签 form:"page" 表示该字段可以从表单数据中获取名为 "page" 的值。
	Key   string `form:"key"`   //标签 form:"key" 表示该字段可以从表单数据中获取名为 "key" 的值。
	Limit int    `form:"limit"` //标签 form:"limit" 表示该字段可以从表单数据中获取名为 "limit" 的值。
	Sort  string `form:"sort"`  //标签 form:"sort" 表示该字段可以从表单数据中获取名为 "sort" 的值。
}
