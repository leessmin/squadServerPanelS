package util

import (
	"fmt"
	"testing"
	"time"
)

// 测试 创建jwt 函数
func TestCreateJWT(t *testing.T) {
	s, err := UtilJWT.CreateJWT("leessmin")
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 3)

	ss, err := UtilJWT.CreateJWT("leessmin")
	if err != nil {
		panic(err)
	}
	fmt.Println("token:", s)
	fmt.Println("token2:", ss)
	fmt.Println("token是否相等:", s == ss)
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImxlZXNzbWluIiwiaXNzIjoic3F1YWRTZXJ2ZXJQYW5lbFNlcnZlciIsInN1YiI6ImxlZXNzbWluIiwiYXVkIjpbImxlZXNzbWluIl0sImV4cCI6MTY4MDEzNzA3NCwibmJmIjoxNjgwMDUwNjc0LCJpYXQiOjE2ODAwNTA2NzQsImp0aSI6IjEifQ.-hxtt1uXhxOIgDp7e2nBjtSL-1_SdUWO-ZIPYIClCvk

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImxlZXNzbWluIiwiaXNzIjoic3F1YWRTZXJ2ZXJQYW5lbFNlcnZlciIsInN1YiI6ImxlZXNzbWluIiwiYXVkIjpbImxlZXNzbWluIl0sImV4cCI6MTY4MDEzNzA3NCwibmJmIjoxNjgwMDUwNjc0LCJpYXQiOjE2ODAwNTA2NzQsImp0aSI6IjEifQ.-hxtt1uXhxOIgDp7e2nBjtSL-1_SdUWO-ZIPYIClCvk
