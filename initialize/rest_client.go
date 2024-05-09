package initialize

import (
	"crypto/tls"
	"dsvm/global"
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"time"
)

func RestClient() {
	restCfg := global.AppConfig.RestClient
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		//InsecureSkipVerify用来控制客户端是否证书和服务器主机名。如果设置为true, 则不会校验证书以及证书中的主机名和服务器主机名是否一致。
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext:     (getDialer()).DialContext,
	}
	// 客户端默认超时 5秒
	timeout := time.Second * 3
	if len(restCfg.Timeout) > 0 {
		d4, err := time.ParseDuration(restCfg.Timeout)
		if err == nil {
			timeout = d4
		}
	}
	resty1 := resty.New().
		SetHeader("Content-Type", "application/json").
		SetDisableWarn(true).
		SetTimeout(timeout).
		SetTransport(tr)

	retryCfg := restCfg.Retry
	resty1.SetRetryCount(retryCfg.Count)
	if len(retryCfg.WaitTime) > 0 {
		d5, err := time.ParseDuration(retryCfg.WaitTime)
		if err == nil {
			resty1.SetRetryWaitTime(d5)
		}
	}
	if len(retryCfg.MaxWaitTime) > 0 {
		d6, err := time.ParseDuration(retryCfg.MaxWaitTime)
		if err == nil {
			resty1.SetRetryMaxWaitTime(d6)
		}
	}
	global.RestClient = resty1
}

func getDialer() *net.Dialer {
	dialerCfg := global.AppConfig.RestClient.Dialer
	dialer := &net.Dialer{}
	if len(dialerCfg.Timeout) > 0 {
		d1, err := time.ParseDuration(dialerCfg.Timeout)
		if err == nil {
			dialer.Timeout = d1
		}
	}
	if len(dialerCfg.FallbackDelay) > 0 {
		d2, err := time.ParseDuration(dialerCfg.FallbackDelay)
		if err == nil {
			dialer.FallbackDelay = d2
		}
	}
	if len(dialerCfg.KeepAlive) > 0 {
		d3, err := time.ParseDuration(dialerCfg.KeepAlive)
		if err == nil {
			dialer.KeepAlive = d3
		}
	}
	return dialer
}
