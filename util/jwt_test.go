package util

import (
	"fmt"
	"testing"
	"time"
)

// 测试 创建jwt 函数
func TestCreateJWT(t *testing.T) {
	s, err := UtilJWT.CreateJWT("leessmin",time.Now())
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 3)

	ss, err := UtilJWT.CreateJWT("leessmin",time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println("token:", s)
	fmt.Println("token2:", ss)
	fmt.Println("token是否相等:", s == ss)
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImxlZXNzbWluIiwiaXNzIjoic3F1YWRTZXJ2ZXJQYW5lbFNlcnZlciIsInN1YiI6ImxlZXNzbWluIiwiYXVkIjpbImxlZXNzbWluIl0sImV4cCI6MTY4MDEzNzA3NCwibmJmIjoxNjgwMDUwNjc0LCJpYXQiOjE2ODAwNTA2NzQsImp0aSI6IjEifQ.-hxtt1uXhxOIgDp7e2nBjtSL-1_SdUWO-ZIPYIClCvk

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImxlZXNzbWluIiwiaXNzIjoic3F1YWRTZXJ2ZXJQYW5lbFNlcnZlciIsInN1YiI6ImxlZXNzbWluIiwiYXVkIjpbImxlZXNzbWluIl0sImV4cCI6MTY4MDEzNzA3NCwibmJmIjoxNjgwMDUwNjc0LCJpYXQiOjE2ODAwNTA2NzQsImp0aSI6IjEifQ.-hxtt1uXhxOIgDp7e2nBjtSL-1_SdUWO-ZIPYIClCvk

// 测试 验证token 函数
func TestVerifyToken(t *testing.T) {
	c := UtilJWT.VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFkbWluIiwiT3BfdGltZSI6IjIwMjMtMDMtMzBUMDc6MzI6MDAtMDg6MDAiLCJpc3MiOiJzcXVhZFNlcnZlclBhbmVsU2VydmVyIiwic3ViIjoiYWRtaW4iLCJhdWQiOlsiYWRtaW4iXSwiZXhwIjoxNjgwMjQ5ODE4LCJuYmYiOjE2ODAxNjM0MTgsImlhdCI6MTY4MDE2MzQxOCwianRpIjoiMSJ9.5wkQlSrnQoKmgT9qocuF3mWYy-yhlE1lvt4tnwWHuFs")

	fmt.Println(c)
}
