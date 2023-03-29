package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCaptchaHandle(t *testing.T) {
	Login.CaptchaHandle(&gin.Context{})
}
