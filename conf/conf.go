package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ResourceConf struct {
	Mysql MysqlConf
	// Redis map[string]redis.RedisConf  `yaml:"redis"`
	// Rmq   map[string]rmq.ClientConfig `yaml:"rmq"`
}

type MysqlConf struct {
	Database        string `yaml:"database"`
	Addr            string `yaml:"addr"`
	Port            string `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Network         string `yaml:"network"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int64  `yaml:"connMaxLifeTime"`
}

var (
	RConf ResourceConf
)

func InitConf() {
	configFile, err := ioutil.ReadFile("conf/resource.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configFile, &RConf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("conf:%v", RConf)
}
