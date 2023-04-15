package route

import (
	_ "SSPS/router/logger"
	_ "SSPS/router/login"
	_ "SSPS/router/panel_config"
	_ "SSPS/router/proxy"
	_ "SSPS/router/squad_admin_group"
	_ "SSPS/router/squad_admin_user"
	_ "SSPS/router/squad_bans"
	_ "SSPS/router/squad_day_msg"
	_ "SSPS/router/squad_license"
	_ "SSPS/router/squad_map"
	_ "SSPS/router/squad_server"
	_ "SSPS/router/squad_server_msg"
	_ "SSPS/router/system"
)

// 路由注册
// 调用模块 模块执行 init 即可注册路由
