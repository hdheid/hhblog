package article_api

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/models/ctype"
	"gvb_server/utils/jwts"
	"math/rand"
	"strings"
	"time"
)

type ArticleRequest struct {
	Title    string      `json:"title" binding:"required" msg:"文章标题必填"`              // 文章标题
	Abstract string      `json:"abstract"`                                           // 文章简介
	Content  string      `json:"content,omit(list)" binding:"required" msg:"文章内容必填"` // 文章内容
	Category string      `json:"category"`                                           // 文章分类
	Source   string      `json:"source"`                                             // 文章来源
	Link     string      `json:"link"`                                               // 原文链接
	BannerID uint        `json:"banner_id"`                                          // 文章封面id
	Tags     ctype.Array `json:"tags"`                                               // 文章标签
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	userNickName := claims.NickName

	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		common.FailWithError(err, &cr, c)
		return
	}

	//这里需要校验 content，防止 xss 攻击
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	//是否有script恶意标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		//表示有恶意script
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}

	if cr.Abstract == "" {
		//这里就从 content 里面去选择 30 个字符
		/*
			在 go 中，[]string 类型里面字符和汉字的字节数是不一样的，可以转化成rune切片，这个时候就是一样的
		*/

		//汉字截取需要转化一下
		abs := []rune(doc.Text()) //这里需要过滤掉类似于 # 这样的符号
		//特判文章长度小于100
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs) //所有
		}
	}

	//如果没有选择封面，就随机选择一张
	if cr.BannerID == 0 {
		var bannerIdList []uint
		global.DB.Model(&models.BannerModel{}).Select("id").Scan(&bannerIdList)
		if len(bannerIdList) == 0 {
			common.FailWithMessage("没有图片！", c)
			return
		}

		rand.Seed(time.Now().UnixNano())
		cr.BannerID = bannerIdList[rand.Intn(len(bannerIdList))]
	}

	//查询banner_id的banner_url
	var bannerurl string
	err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerurl).Error
	if err != nil {
		common.FailWithMessage("图片不存在！", c)
		global.Log.Debug("图片不存在！")
		return
	}

	//查询用户头像
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("Avatar").Scan(&avatar).Error
	if err != nil {
		common.FailWithMessage("用户不存在！", c)
		global.Log.Debug("用户不存在！")
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	//创建一个es数据
	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerurl,
		Tags:         cr.Tags,
	}

	err = article.Create()
	if err != nil {
		common.FailWithMessage("文章创建失败！", c)
		global.Log.Debug("文章创建失败！")
		return
	}

	common.OKWithMessage("文章创建成功！", c)
}
