package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Name           string             `yaml:"Name,default=dsvm" json:"Name,default=dsvm"`
	Version        string             `yaml:"Version,default=v1.1.1" json:"Version,default=v1.1.1"`
	RpcServerConf  zrpc.RpcServerConf `yaml:"RpcServerConf" json:"RpcServerConf"`
	RestServerConf rest.RestConf      `yaml:"RestServerConf" json:"RestServerConf"`
	LogxConf       logx.LogConf       `yaml:"LogxConf" json:"LogxConf"`
}
