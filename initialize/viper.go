package initialize

import (
	"dsvm/conf"
	"dsvm/global"
	"errors"
	"flag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
)

func Viper() {
	viper.SetDefault("appName", "dsvm")
	viper.SetDefault("version", "v1.1.1")
	viper.SetDefault("httpPort", 12888)
	viper.SetDefault("grpcPort", 12886)
	viper.SetDefault(
		"restClient",
		conf.RestClientConfig{
			Dialer: conf.DialerConfig{
				Timeout:       "5s",
				FallbackDelay: "1s",
				KeepAlive:     "60s",
			},
			Retry: conf.RetryConfig{
				Count:       3,
				WaitTime:    "1s",
				MaxWaitTime: "10s",
			},
			Timeout: "3s",
		},
	)
	var config string
	flag.StringVar(&config, "c", "config.yaml", "choose config file.")
	flag.Parse()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(config)
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			log.Panicf("viper ReadInConfig error : %v \n", err)
		}
	}
	if err := viper.Unmarshal(&global.AppConfig); err != nil {
		log.Panicf("viper Unmarshal error : %v \n", err.Error())
	}
	out, _ := yaml.Marshal(global.AppConfig)
	log.Printf("\n%s\n", string(out))
}
