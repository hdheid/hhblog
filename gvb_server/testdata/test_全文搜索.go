package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"strings"
)

/*
关于 # 开头的一定是标题 ？
- 那么在代码块里面的#如果也是在开头呢 ？ 需要判断一下是否在代码块里面，如果用户只写了开头三个点没写结尾的，那么下面的内容永远都会渲染在代码块中，和下面的 isCode 逻辑保持一致
- 或者说写的不够标准，比如#不打空格：可以使用正则表达式来判断：#{1,6} *? 之类的
- 有空格怎么办？ 标题前面全是空格，那么这坨空格也会被识别到，所以应该先去掉空格后再看有没有内容，如果有就加进来
- 我不需要标题列表里面的#，还有正文里面的一些图片链接之类的

- headList 和 bodyList 长度一定要相同，为什么？ 去掉body的空值处理
*/

func main() {
	var _ = "## 环境搭建\n\n拉取镜像\n\n```Python\ndocker pull elasticsearch:7.12.0\n```\n\n\n\n创建docker容器挂在的目录：\n\n```Python\nmkdir -p /opt/elasticsearch/config & mkdir -p /opt/elasticsearch/data & mkdir -p /opt/elasticsearch/plugins\n\nchmod 777 /opt/elasticsearch/data\n\n```\n\n配置文件\n\n```Python\necho \"http.host: 0.0.0.0\" >> /opt/elasticsearch/config/elasticsearch.yml\n```\n\n\n\n创建容器\n\n```Python\n# linux\ndocker run --name es -p 9200:9200  -p 9300:9300 -e \"discovery.type=single-node\" -e ES_JAVA_OPTS=\"-Xms84m -Xmx512m\" -v /opt/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /opt/elasticsearch/data:/usr/share/elasticsearch/data -v /opt/elasticsearch/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0\n```\n\n\n\n访问ip:9200能看到东西\n\n![](http://python.fengfengzhidao.com/pic/20230129212040.png)\n\n就说明安装成功了\n\n\n\n浏览器可以下载一个 `Multi Elasticsearch Head` es插件\n\n\n\n第三方库\n\n```Go\ngithub.com/olivere/elastic/v7\n```\n\n## es连接\n\n```Go\nfunc EsConnect() *elastic.Client  {\n  var err error\n  sniffOpt := elastic.SetSniff(false)\n  host := \"http://127.0.0.1:9200\"\n  c, err := elastic.NewClient(\n    elastic.SetURL(host),\n    sniffOpt,\n    elastic.SetBasicAuth(\"\", \"\"),\n  )\n  if err != nil {\n    logrus.Fatalf(\"es连接失败 %s\", err.Error())\n  }\n  return c\n}\n```"
}

type SearchData struct {
	Body  string `json:"body"`  //正文
	Slug  string `json:"slug"`  //跳转地址
	Title string `json:"title"` //标题
}

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
