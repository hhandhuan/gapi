package conf

import (
	"flag"
	"zhengze/pkg/utils"
)

type Jwt struct {
	Issuer string `yaml:"issuer"`
	Ttl    int    `yaml:"ttl"`
	Secret string `yaml:"secret"`
}

type Mysql struct {
	Dns             string `yaml:"dns"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxIdleTime int    `yaml:"connMaxIdleTime"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Pass string `yaml:"pass"`
	DB   int    `yaml:"db"`
}

type System struct {
	Env  string `yaml:"env"`
	Addr string `yaml:"addr"`
}

type Config struct {
	System *System `yaml:"system"`
	Mysql  *Mysql  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
	Jwt    *Jwt    `yaml:"jwt"`
}

var (
	config = new(Config)
)

var (
	cPath = flag.String("cfg", "./config/config.yaml", "config file path")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*cPath, config)
}

func GetConfig() *Config {
	return config
}
