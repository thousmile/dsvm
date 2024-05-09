package main

import (
	"dsvm/core"
	"dsvm/initialize"
	"embed"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	// 初始化 viper config
	initialize.Viper()

	// 初始化 rest client
	initialize.RestClient()

	// 初始化 数据源
	initialize.DataSource()
	// 初始化 goose
	goose.SetBaseFS(embedMigrations)

	// 启动 http server
	core.RunHttpServer()
}
