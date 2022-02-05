package config

type Config struct {
	Mode      string          `yaml:"mode" json:"mode"`
	Ticker    TickerConfig    `yaml:"ticker" json:"ticker"`
	ApiServer ApiServerConfig `yaml:"apiServer" json:"apiServer"`
}

type TickerConfig struct {
	Interval int32 `yaml:"interval" json:"interval"`
}

type ApiServerConfig struct {
	Addr    string `yaml:"addr" json:"addr"`
	Timeout int32  `yaml:"timeout" json:"timeout"`
}
