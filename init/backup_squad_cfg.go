package init

import (
	"SSPS/config"
	"SSPS/util"
)

// 备份 squad 服务器的配置
func BackupSquadCfg() {
	// 判断是否存在 squad 服务器 目录
	if len(config.CreatePanelConf().GameServePath) <= 0 {
		// 不存在目录 退出
		return
	}

	// 判断是否有目录
	b, _ := util.CreateReadWrite().IsDir("./backCfg")

	// 没有目录创建目录  并备份文件
	if !b {
		// 创建目录
		util.CreateReadWrite().CreateDir("./backCfg")
		// 备份文件
		backup()
	}
}

func backup() {
	// 备份配置文件
	err := util.CreateReadWrite().CopyFile("LayerRotation.cfg", "./backCfg/")
	if err != nil {
		panic(err)
	}

	err = util.CreateReadWrite().CopyFile("LevelRotation.cfg", "./backCfg/")
	if err != nil {
		panic(err)
	}
}
