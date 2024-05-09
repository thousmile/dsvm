package conf

type AppServerConfig struct {
	// 服务名称
	AppName string `mapstructure:"appName" yaml:"appName" json:"appName"`

	// http 服务端口
	Version string `mapstructure:"version" yaml:"version" json:"version"`

	// http 服务端口
	HttpPort int `mapstructure:"httpPort" yaml:"httpPort" json:"httpPort"`

	// grpc 服务端口
	GrpcPort int `mapstructure:"grpcPort" yaml:"grpcPort" json:"grpcPort"`

	// rest client
	RestClient RestClientConfig `mapstructure:"restClient" json:"restClient" yaml:"restClient"`

	// DataSources
	DataSources map[string]DataSourceConfig `mapstructure:"dataSources" json:"dataSources" yaml:"dataSources"`
}
