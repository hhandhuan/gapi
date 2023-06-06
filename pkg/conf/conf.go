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
	Env              string `yaml:"env"`
	Addr             string `yaml:"addr"`
	ShutdownWaitTime int    `yaml:"shutdownWaitTime"`
}

type Logger struct {
	Path       string `yaml:"path"`
	Level      int    `yaml:"level"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
	Compress   bool   `yaml:"compress"`
}

type Config struct {
	System *System `yaml:"system"`
	Mysql  *Mysql  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
	Jwt    *Jwt    `yaml:"jwt"`
	Logger *Logger `yaml:"logger"`
}

var (
	conf = new(Config)
	path = flag.String("cfg", "./config/config.yaml", "config file path")
)

func Initialize() *Config {
	flag.Parse()
	var buf []byte
	buf, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	return conf
}

func GetConfig() *Config {
	return conf
}
