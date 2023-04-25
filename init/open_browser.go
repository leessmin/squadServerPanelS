package init

import (
	"SSPS/config"
	"SSPS/util"
	"fmt"
)

// 打开默认浏览器
func init() {
	util.OpenBrowser(fmt.Sprintf("http://%v:%v/index", config.CreatePanelConf().ServerIp, config.CreatePanelConf().PanelPort))
}
