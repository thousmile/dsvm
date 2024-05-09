package conf

type RestClientConfig struct {
	Dialer DialerConfig `mapstructure:"dialer" json:"dialer" yaml:"dialer"`

	Retry RetryConfig `mapstructure:"retry" json:"retry" yaml:"retry"`

	Timeout string `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}

type DialerConfig struct {
	Timeout string `mapstructure:"timeout" json:"timeout" yaml:"timeout"`

	FallbackDelay string `mapstructure:"fallbackDelay" json:"fallbackDelay" yaml:"fallbackDelay"`

	KeepAlive string `mapstructure:"keepAlive" json:"keepAlive" yaml:"keepAlive"`
}

type RetryConfig struct {
	Count int `mapstructure:"count" json:"count" yaml:"count"`

	WaitTime string `mapstructure:"waitTime" json:"waitTime" yaml:"waitTime"`

	MaxWaitTime string `mapstructure:"maxWaitTime" json:"maxWaitTime" yaml:"maxWaitTime"`
}
