package util

import (
	"bytes"
	"encoding/base64"
	"fmt"

	_ "image/png"

	"github.com/dchest/captcha"
)

// 验证码 结构体
type Captcha struct {
	// 验证码的id
	Id string `json:"id"`
	// 验证码的图片
	Image string `json:"image"`
}

// 创建验证码 w70px h35px
func CreateCaptcha(w, h int) *Captcha {
	// 实例化 验证码结构体
	t_captcha := Captcha{}

	// 创建验证码   长度为4
	t_captcha.Id = captcha.NewLen(4)

	t_captcha.Image = createBase64(t_captcha.Id, w, h)

	// 返回验证码图片 结构体
	return &t_captcha
}

// 验证 验证码
func VerifyCaptcha(id string, digits string) bool {
	return captcha.VerifyString(id, digits)
}

// 通过 验证码的id 生成 base64编码的图片
func createBase64(id string, w int, h int) string {
	// 储存图片 buffer
	var contentImg bytes.Buffer

	// 生成图片 写入contentImg
	err := captcha.WriteImage(&contentImg, id, w, h)
	if err != nil {
		GetError().ServerError(fmt.Sprint("生成验证码出错，err:", err))
	}

	// 创建一个 byte 储存图片编码后的结果
	dist := make([]byte, 5000)
	// base64编码
	base64.StdEncoding.Encode(dist, contentImg.Bytes())

	// 去除 \u0000
	dist = bytes.Trim(dist, "\u0000")

	// 生成完毕 返回base64编码
	return fmt.Sprint("data:image/png;base64,", string(dist))
}
