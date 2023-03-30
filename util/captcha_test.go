package util

import (
	"fmt"
	_ "image/png"
	"testing"

	"github.com/dchest/captcha"
)

// 测试创建验证码函数
func TestCreateBase64(t *testing.T) {
	id := captcha.NewLen(4)

	s := createBase64(id, 70, 35)

	fmt.Println(s)
}
