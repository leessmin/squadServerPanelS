package util

import (
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
	jwt.RegisteredClaims
}

// 创建jwt
func (u uJWT) CreateJWT(username string) (string, error) {
	// 实例化 签发 jwt 的结构体
	claims := customClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			// TODO:暂定24小时
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
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
