package config

type Config struct {
	Mode       string           `yaml:"mode" json:"mode"`
	Http       HttpConfig       `yaml:"http" json:"http"`
	UserServer UserServerConfig `yaml:"userServer" json:"userServer"`
	Tracer     TracerConfig     `yaml:"tracer" json:"tracer"`
}

type HttpConfig struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type UserServerConfig struct {
	Addr    string `yaml:"addr" json:"addr"`
	Timeout int32  `yaml:"timeout" json:"timeout"`
}

type TracerConfig struct {
	Enabled     bool   `yaml:"enabled" json:"enabled"`
	ServiceName string `yaml:"serviceName" json:"serviceName"`
	Endpoint    string `yaml:"endpoint" json:"endpoint"`
}
