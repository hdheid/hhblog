package es_ser

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"gvb_server/global"
	"gvb_server/models"
	"strings"
)

type SearchData struct {
	Key   string `json:"key"`   //关键字
	Body  string `json:"body"`  //正文
	Slug  string `json:"slug"`  //跳转地址
	Title string `json:"title"` //标题
}

// GetSearchIndexDataByContent 对文章进行处理，生成标准的格式，方便同步
func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n") //按行进行分割
	var headList, bodyList []string
	var body string

	//加入文章标题
	headList = append(headList, GetHeader(title))

	var isCode = false //默认不在代码块里面，也就是```xxx```中
	for _, s := range dataList {
		if strings.HasPrefix(s, "```") {
			isCode = !isCode //当第一次碰到的时候，表示已经在代码块里面了，所以这里用的是取反操作，下一次碰到就结束代码块
		}

		if strings.HasPrefix(s, "#") && !isCode {
			headList = append(headList, GetHeader(s))
			//if strings.TrimSpace(body) != "" { //这里处理了空值问题
			bodyList = append(bodyList, GetBody(body)) //遇到标题后，将上面一个标题到这个标题之间的正文内容加到 bodyList 中
			body = ""                                  //将 body 清空
			//}
			continue
		}

		body += s //将正文内容加起来
	}
	bodyList = append(bodyList, GetBody(body)) //将最后一个标题后面的内容添加到 bodyList 里面

	lenth := len(headList)
	for i := 0; i < lenth; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + Getslug(headList[i]),
			Key:   id,
		})
	}

	return
}

func GetHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, "_", "")
	return head
}

func GetBody(body string) string {
	bodyHtml := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(bodyHtml)))
	return doc.Text() //去掉了图片链接
}

func Getslug(slug string) string {
	return "#" + slug
}

// AsyncArticleByFullText 同步文章数据到全文搜索
func AsyncArticleByFullText(id, title, content string) error {
	indexList := GetSearchIndexDataByContent(id, title, content)

	//批量添加
	bulkService := global.Client.Bulk()

	for _, data := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(data)
		bulkService.Add(req)
	}

	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return err
	}

	global.Log.Infof("%s 添加成功！共 %d 条", title, len(result.Succeeded()))

	return nil
}

// DeleteFullTextByArticleId 从全文搜索中删除数据
func DeleteFullTextByArticleId(id string) {
	boolSearch := elastic.NewTermQuery("key", id)

	res, _ := global.Client.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())

	global.Log.Infof("删除成功！共删除 %d 条", res.Deleted)
}
