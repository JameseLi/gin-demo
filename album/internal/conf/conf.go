package conf

import "time"

type BootStrap struct {
	Server     Server     `yaml:"server"`
	Datasource Datasource `yaml:"datasource"`
	Redis      Redis      `yaml:"redis"`
	Jwt        Jwt        `yaml:"jwt"`
}

type Server struct {
	Http Http `yaml:"http"`
}

type Http struct {
	Port         int32         `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type Datasource struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}

type Redis struct {
	Addr string `yaml:"addr"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}

func NewBootStrap() *BootStrap {
	return &BootStrap{}
}
