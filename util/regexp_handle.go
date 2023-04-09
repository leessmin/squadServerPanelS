package util

import (
	"strings"
	"sync"

	"github.com/dlclark/regexp2"
)

// 正则表达式结构体
type Regexp3 struct{}

var onceRegexp sync.Once

var regexp3 *Regexp3

// 创建正则表达式
func CreateRegexp() *Regexp3 {
	onceRegexp.Do(func() {
		regexp3 = &Regexp3{}
	})

	return regexp3
}

// 验证 字符串是否符合 正则表达式
// 符合 true   不符合 false
func (r *Regexp3) VerifyStr(pattern, str string) bool {
	// 判断是否符合正则表达式
	re := regexp2.MustCompile(pattern, 0)
	// 判断 字符串是否符合 正则表达式
	isOk, _ := re.MatchString(str)

	return isOk
}

// 根据正则表达式 提取字符串
// 如果存在 可以提取的字符串  bool 类型返回true 否 返回false
func (r *Regexp3) FindString(pattern, str string) ([]string, bool) {
	re := regexp2.MustCompile(pattern, 0)

	m, _ := re.FindStringMatch(str)
	if m == nil {
		return nil, false
	}

	// 获取提取到的字符串组
	gps := m.Groups()

	var strArr []string
	// 遍历拿到每个字符串
	for _, v := range gps {
		strArr = append(strArr, strings.TrimSpace(v.String()))
	}

	return strArr, true
}
