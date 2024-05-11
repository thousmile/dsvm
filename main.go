package main

import (
	"dsvm/core"
	"dsvm/initialize"
	"dsvm/utils"
	"github.com/pressly/goose/v3"
	"os"
)

func main() {
	// 初始化 viper config
	initialize.Viper()

	// 初始化 rest client
	initialize.RestClient()

	// 初始化 数据源
	initialize.DataSource()

	// 创建目录
	_ = utils.CreateDir(core.MigrationsDir)
	dir, _ := os.ReadDir(core.MigrationsDir)
	core.WriteMigrationsIndex(len(dir))

	// 初始化 goose
	goose.SetBaseFS(os.DirFS(core.MigrationsDir))

	// 启动 http server
	core.RunHttpServer()
}
