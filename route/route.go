package route

import (
	_ "SSPS/router/login"
	_ "SSPS/router/proxy"
	_ "SSPS/router/squad_admin_group"
	_ "SSPS/router/squad_admin_user"
	_ "SSPS/router/system"
	_ "SSPS/router/squad_bans"
	_ "SSPS/router/squad_server"
	_ "SSPS/router/squad_map"
)

// 路由注册
// 调用模块 模块执行 init 即可注册路由
