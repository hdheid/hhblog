package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// ImageNameListView 获取图片部分数据
// @tags 图片管理
// @Summary 查询部分图片信息
// @Description 查询所有的图片的ID、地址、名字
// @Router /api/images_names [get]
// @Produce json
// @Success 200 {object} common.Response{data=[]ImageResponse}
func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imageNameList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageNameList)
	common.OKWithData(imageNameList, c)
}
