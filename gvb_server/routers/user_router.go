package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

// 如果在其他路由里面也需要用 session 的话，就可以定义在 enter 里面，目前只用在 user 里卖，所以暂时定义在这里
var store = cookie.NewStore([]byte("WszidjnchHusoJSHdiqksds"))

func UserRouter(r *gin.RouterGroup) {
	userApi := api.ApiGroupApp.UserApi

	r.Use(sessions.Sessions("sessionid", store)) //session 的使用，后期可以考虑自己写一个session的中间件，更方便控制

	r.POST("/email_login", userApi.EmailLoginView)
	r.POST("/users", userApi.UserCreateView)
	r.GET("/users", middleware.JwtAuth(), userApi.UserListView)
	r.PUT("/user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	r.PUT("/user_password", middleware.JwtAuth(), userApi.UserUpdatePassword)
	r.POST("/user_logout", middleware.JwtAuth(), userApi.UserLogoutView)
	r.DELETE("/user", middleware.JwtAuth(), userApi.UserRemoveView)
	r.POST("/user_bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)

}
