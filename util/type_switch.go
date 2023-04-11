package util

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// 类型强制转换

// 类型强制转换 结构体
type typeSwitch struct{}

var (
	t        *typeSwitch
	typeOnce sync.Once
)

func CreateTypeSwitch() *typeSwitch {
	typeOnce.Do(func() {
		t = &typeSwitch{}
	})
	return t
}

// 字符串转bool
func (t *typeSwitch) StringToBool(str string) (bool, error) {
	str = strings.TrimSpace(str)
	if str == "0" || str == "1" {
		return false, fmt.Errorf("有0和1")
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}

	return b, nil
}

// 字符串转int
func (t *typeSwitch) StringToInt(str string) (int, error) {
	str = strings.TrimSpace(str)
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}
