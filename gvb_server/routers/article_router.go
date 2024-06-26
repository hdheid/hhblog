package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func ArticleRouter(r *gin.RouterGroup) {
	articleApi := api.ApiGroupApp.ArticleApi
	r.POST("/articles", middleware.JwtAdmin(), articleApi.ArticleCreateView)
	r.GET("/articles", articleApi.ArticleListView)
	r.GET("/articles/:id", articleApi.ArticleDetailView)
	r.GET("/articles/detail", articleApi.ArticleDetaiByTitlelView)
	r.GET("/articles/calendar", articleApi.ArticleCalendarView)
	r.GET("/articles/tags", articleApi.ArticleTagListView)
	r.PUT("/articles", articleApi.ArticleUpdateView)
	r.DELETE("/articles", articleApi.ArticleRemoveView)
	r.POST("/articles/collects", middleware.JwtAuth(), articleApi.ArticleCollCreateView)
	r.GET("/articles/collects", middleware.JwtAuth(), articleApi.ArticleCollListView)
	r.DELETE("/articles/collects", middleware.JwtAuth(), articleApi.ArticleCollBatchRemoveView)
	r.GET("/articles/text", articleApi.FullTextSearchView) //全文搜索接口
}
