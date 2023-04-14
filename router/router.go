package router

import (
	"SSPS/middleware"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

var APIGroup *gin.RouterGroup

var BasicAuth *gin.RouterGroup

func init() {
	// 启动gin
	Router = gin.New()
	{
		// 挂载中间件  全局
		Router.Use(middleware.ErrorMiddlewareHandle())

		// 挂载 解决跨域插件  全局
		Router.Use(middleware.CorsMiddlewareHandle())
	}

	// 创建api路由组
	APIGroup = Router.Group("/api")

	// 需要鉴权的路由组 鉴权
	BasicAuth = APIGroup.Group("/BA")
	{
		// 挂载鉴权插件
		BasicAuth.Use(middleware.AuthMiddlewareHandle())
	}
}
