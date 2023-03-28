package router

import "github.com/gin-gonic/gin"

var Router *gin.Engine

var APIGroup *gin.RouterGroup

func init() {
	// 启动gin
	Router = gin.Default()

	// 创建api路由组
	APIGroup = Router.Group("/api")
}
