package model

import "database/sql"

type GooseDataSource struct {
	Dialect         string  `json:"dialect" form:"dialect"`
	CreateSchemaSql string  `yaml:"createSchemaSql" json:"createSchemaSql"`
	SwitchSchemaSql string  `yaml:"switchSchemaSql" json:"switchSchemaSql"`
	Db              *sql.DB `json:"db" form:"db"`
}
