package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var UtilJWT uJWT

// jwt工具包
type uJWT struct {
	// 签发秘钥
	signingKey []byte
}

func init() {
	// 实例 utilJWT
	UtilJWT = uJWT{
		signingKey: []byte("leessmin"),
	}
}

// 签发 JWT 结构体
type customClaims struct {
	// 用户名
	Username string
	Op_time  int64
	jwt.RegisteredClaims
}

// 创建jwt
func (u uJWT) CreateJWT(username string, op_time int64) (string, error) {
	// 实例化 签发 jwt 的结构体
	claims := customClaims{
		Username: username,
		Op_time:  op_time,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			// TODO:暂定jwt过期时间为1年
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(1, 0, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "squadServerPanelServer",
			Subject:   username,
			// TODO: ID暂且拟定为1
			ID:       "1",
			Audience: []string{username},
		},
	}

	// 创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成 token string
	tokenString, err := token.SignedString(u.signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 验证token
func (u uJWT) VerifyToken(tokenString string) *customClaims {

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		return u.signingKey, nil
	})

	if err != nil {
		GetError().ForbiddenError(fmt.Sprint("token无效,err:", err))
	}

	// 将token转为 claims
	claims, ok := token.Claims.(*customClaims)

	if !(ok && token.Valid) {
		GetError().ForbiddenError("token无效")
	}

	// 返回 JWT 结构体
	return claims
}
