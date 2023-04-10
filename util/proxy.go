package util

import (
	"io"
	"net/http"
	"net/url"
	"sync"
)

// 代理http请求

type httpProxy struct{}

var (
	Proxy    *httpProxy
	httpOnce sync.Once
)

func CreateHttpProxy() *httpProxy {
	httpOnce.Do(func() {
		Proxy = &httpProxy{}
	})
	return Proxy
}

// 代理get请求
func (hp *httpProxy) ProxyGet(url string, query url.Values) []byte {
	// 创建http client
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// 拼接query参数
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer func(resp *http.Response) {
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}
