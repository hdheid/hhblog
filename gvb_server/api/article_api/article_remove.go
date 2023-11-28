package article_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	client := global.Client
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	bulkService := client.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")

	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}

	res, err := bulkService.Do(context.Background())
	if err != nil {

		global.Log.Error(err)
		common.FailWithMessage("删除失败！", c)
	}

	common.OKWithMessage(fmt.Sprintf("成功删除文章 %d 篇！", len(res.Succeeded())), c)
}
