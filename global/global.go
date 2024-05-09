package global

import (
	"dsvm/conf"
	"dsvm/model"
	"github.com/go-resty/resty/v2"
	"sync"
)

var (
	AppConfig *conf.AppServerConfig

	RestClient *resty.Client

	DataSources map[string]*model.GooseDataSource

	lock sync.RWMutex
)
