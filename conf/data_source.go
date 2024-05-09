package conf

type DataSourceConfig struct {
	Dialect string `json:"dialect" form:"dialect"`
	// 创建 Schema 的sql语句
	CreateSchemaSql string `yaml:"createSchemaSql" json:"createSchemaSql"`
	// 切换 Schema 的sql语句
	SwitchSchemaSql string `yaml:"switchSchemaSql" json:"switchSchemaSql"`
	// 地址
	Addr string `yaml:"addr" json:"addr"`
	// 端口
	Port int `yaml:"port" json:"port"`
	// 地址
	DbName string `yaml:"dbName" json:"dbName"`
	// 用户名，默认: root
	Username string `yaml:"username" json:"username"`
	// 密码，默认: root
	Password string `yaml:"password" json:"password"`
}
