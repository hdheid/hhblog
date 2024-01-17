package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

/*
在给结构体字段添加标签时，binding 是 Gin 框架中的一个验证标签。它可以用于验证结构体字段的值是否符合指定的条件。

在你提供的代码中，binding 标签被用于对结构体字段进行验证。

binding:"required"：这个标签指示该字段是必需的，即在进行请求绑定时，该字段的值不能为零值或空值。
binding:"url"：这个标签指示该字段的值必须是一个合法的 URL。
当使用 Gin 框架处理 HTTP 请求时，你可以使用这些验证标签来确保请求中的数据满足特定的条件。例如，对于 Title 字段，binding:"required" 标签要求请求中必须包含一个非空的 title 值；对于 Href 和 Images 字段，binding:"required,url" 标签要求请求中必须包含非空值且符合 URL 格式。

如果请求中的字段不满足标签中指定的条件，Gin 框架将返回一个验证错误，并中断请求处理。

通过使用 binding 标签，你可以方便地进行请求数据的验证，以确保它们符合你的期望。这有助于减少错误数据的处理并提高代码的健壮性。
*/

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`          // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"输入链接非法" structs:"href"`       // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"输入图片地址非法" structs:"images"` // 图片
	IsShow bool   `json:"is_show"  msg:"请选择是否展示" structs:"is_show"`                      // 是否展示
}

// AdvertCreateView 广告的添加操作
// @tags 广告管理
// @Summary 创建广告
// @Description 关于广告的创建的API，拥有判重逻辑
// @Param data body AdvertRequest	true  "表示多个参数"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} common.Response{}
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	//参数绑定
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithError(err, &cr, c)
		return
	}

	/*
		重复判断：
			1. 使用唯一索引，某个字段在入库的时候重复，那么 create 会报错
			2. 使用钩子函数，在入库前进行判重处理
			3. 简单处理，如下直接写一个判重的逻辑
	*/
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		common.FailWithMessage("广告已存在，请勿重复添加！", c)
		return
	}

	//添加进入数据库
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error

	if err != nil {
		global.Log.Error("创建广告失败：", err)
		common.FailWithMessage("添加广告失败！", c)
		return
	}

	common.OKWithMessage("添加广告成功！", c)

}
