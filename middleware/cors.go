package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨域

func CorsMiddlewareHandle() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"*"},
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}
