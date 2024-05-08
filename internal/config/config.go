package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RpcServerConf  zrpc.RpcServerConf `yaml:"RpcServerConf" json:"RpcServerConf"`
	RestServerConf rest.RestConf      `yaml:"RestServerConf" json:"RestServerConf"`
	LogxConf       logx.LogConf       `yaml:"LogxConf" json:"LogxConf"`
}
