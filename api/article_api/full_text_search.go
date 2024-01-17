package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

func (ArticleApi) FullTextSearchView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	boolQuery := elastic.NewBoolQuery() //全文搜
	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body")) //按照key搜
	}

	res, err := global.Client.
		Search(models.FullTextModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")).
		Size(100).
		Do(context.Background())
	if err != nil {
		return
	}

	count := res.Hits.TotalHits.Value
	fullTextList := make([]models.FullTextModel, 0)
	for _, hit := range res.Hits.Hits {
		var model models.FullTextModel
		json.Unmarshal(hit.Source, &model)

		if body, ok := hit.Highlight["body"]; ok {
			model.Body = body[0]
		}

		fullTextList = append(fullTextList, model)
	}

	common.OKWithList(fullTextList, count, c)
}
