package static

import (
	"SSPS/router"
)

// 托管 静态文件
// 显示前端页面

func init() {
	// 挂载静态文件
	router.Router.Static("/index/", "./static/web/")

}
