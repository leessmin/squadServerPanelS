package middleware

import (
	"SSPS/config"
	"SSPS/util"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件

func AuthMiddlewareHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 判断是否传递token
		if len(ctx.GetHeader("Authorization")) <= 0 {
			util.GetError().ParameterError("没有发送token")
		}

		// 获取token
		// 从请求头中 获取token
		tokenString := ctx.GetHeader("Authorization")[7:]

		// 解析token
		claims := util.UtilJWT.VerifyToken(tokenString)

		// 创建读取 auth 配置的实例
		authStruct := config.AuthUser{}
		// 读取配置文件  获取登录账号与密码 与修改时间
		configUser := authStruct.ReadAuthConfig()

		// 判断token是否有效
		if configUser.Username != claims.Username || configUser.Op_time != claims.Op_time {
			util.GetError().ForbiddenError("token失效")
		}

		// 通过 放行
		ctx.Next()
	}
}
