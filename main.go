package main

import (
	_ "SSPS/route"
	"SSPS/router"
)

func main() {

	// 启动gin
	router.Router.Run(":8080")
}
