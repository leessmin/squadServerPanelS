package util

import "github.com/gin-gonic/gin"

// 响应数据的格式封装

// 响应消息结构体
type ResponseMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data gin.H  `json:"data"`
}

// 响应消息 结构体 构造函数
func CreateResponseMsg(code int, msg string, data gin.H) *ResponseMsg {
	return &ResponseMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
