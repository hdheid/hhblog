package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Model(models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList) //修改或者删除的时候，model才需要传指针，否则不传指针也可以

	common.OKWithData(menuNameList, c)
}
