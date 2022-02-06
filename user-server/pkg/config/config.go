package config

type Config struct {
	Mode   string       `yaml:"mode" json:"mode"`
	Http   HttpConfig   `yaml:"http" json:"http"`
	Tracer TracerConfig `yaml:"tracer" json:"tracer"`
}

type HttpConfig struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type TracerConfig struct {
	Enabled     bool   `yaml:"enabled" json:"enabled"`
	ServiceName string `yaml:"serviceName" json:"serviceName"`
	Endpoint    string `yaml:"endpoint" json:"endpoint"`
}
