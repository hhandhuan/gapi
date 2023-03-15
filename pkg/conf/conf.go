package conf

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
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
	var buf []byte
	buf, err := ioutil.ReadFile(*cPath)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	err = yaml.Unmarshal(buf, config)
}

func GetConfig() *Config {
	return config
}
