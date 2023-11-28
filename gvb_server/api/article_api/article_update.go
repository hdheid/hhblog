package article_api

import (
	"context"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/models/ctype"
	"time"
)

type ArticleUpdateRequest struct {
	Title    string      `json:"title"`              // 文章标题
	Abstract string      `json:"abstract"`           // 文章简介
	Content  string      `json:"content,omit(list)"` // 文章内容
	Category string      `json:"category"`           // 文章分类
	Source   string      `json:"source"`             // 文章来源
	Link     string      `json:"link"`               // 原文链接
	BannerID uint        `json:"banner_id"`          // 文章封面id
	Tags     ctype.Array `json:"tags"`               // 文章标签
	ID       string      `json:"id"`
}

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	client := global.Client

	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithError(err, &cr, c)
		return
	}

	var bannerurl string
	if cr.BannerID != 0 {
		err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerurl).Error
		if err != nil || bannerurl == "" { //补充逻辑，因为当查询不到数据的时候err同样为空
			common.FailWithMessage("图片不存在！", c)
			global.Log.Debug("补充逻辑！图片不存在！")
			return
		}
	}

	//bannerurl如果为空不要紧，后面会去掉空值
	article := models.ArticleModel{
		UpdatedAt: time.Now().Format(time.DateTime),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: bannerurl,
		Tags:      cr.Tags,
	}

	//判断该文章是否存在
	err = article.GetDataByID(cr.ID)
	if err != nil {
		global.Log.Error(err)
		common.FailWithMessage("文章不存在！", c)
		return
	}

	//如果tags不为空，这里需要加一个tags入库的逻辑，遍历tas，判断数据库里面有没有，然后添加进去

	maps := structs.Map(&article)
	var dataMap = make(map[string]any)
	//去空值
	for key, val := range maps {
		switch val := val.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		}
		dataMap[key] = val
	}

	_, err = client.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(cr.ID).
		Doc(dataMap).
		Do(context.Background())

	if err != nil {
		global.Log.Errorf("更新失败：%s", err.Error())
		common.FailWithMessage("更新失败！", c)
		return
	}

	common.OKWithMessage("更新成功！", c)
}
