package main

import (
	"SSPS/config"
	_ "SSPS/init"
	_ "SSPS/route"
	"SSPS/router"
	"fmt"
)

func main() {

	// 启动gin
	router.Router.Run(fmt.Sprintf(":%v", config.PanelConf.PanelPort))
}
