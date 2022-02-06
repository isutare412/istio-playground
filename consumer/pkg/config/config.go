package config

type Config struct {
	Mode      string          `yaml:"mode" json:"mode"`
	Ticker    TickerConfig    `yaml:"ticker" json:"ticker"`
	ApiServer ApiServerConfig `yaml:"apiServer" json:"apiServer"`
	Tracer    TracerConfig    `yaml:"tracer" json:"tracer"`
}

type TickerConfig struct {
	Interval int32 `yaml:"interval" json:"interval"`
}

type ApiServerConfig struct {
	Addr    string `yaml:"addr" json:"addr"`
	Timeout int32  `yaml:"timeout" json:"timeout"`
}

type TracerConfig struct {
	Enabled     bool   `yaml:"enabled" json:"enabled"`
	ServiceName string `yaml:"serviceName" json:"serviceName"`
	Endpoint    string `yaml:"endpoint" json:"endpoint"`
}
