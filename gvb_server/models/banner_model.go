package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片的类型， 本地还是七牛，设置默认值为1，表示在本地
}

// BeforeDelete 添加一个钩子函数
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	//如果是存储在本地，就同时还需要去删除本地的存储
	if b.ImageType == ctype.Local {
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error("删除本地图片失败！", err)
			return err
		}
	}
	return nil
}
