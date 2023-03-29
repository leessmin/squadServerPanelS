package router

import (
	"SSPS/middleware"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

var APIGroup *gin.RouterGroup

func init() {
	// 启动gin
	Router = gin.Default()
	{
		// 挂载中间件  全局
		Router.Use(middleware.ErrorMiddlewareHandle())
	}

	// 创建api路由组
	APIGroup = Router.Group("/api")
}
