package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"ebay_dapp_golang/pkg/logger"
	"ebay_dapp_golang/pkg/mysql"
	"ebay_dapp_golang/pkg/redis"

	"gopkg.in/yaml.v2"
)

//ServeConf
type ServeConf struct {
	Listen string `yaml:"listen"`
	Debug  bool   `yaml:"debug"`
	Secret string `yaml:"secret"`
}

type ConfYaml struct {
	Serve ServeConf               `yaml:"serve"`
	MySQL mysql.MultipleMysqlConf `yaml:"mysql"`
	Redis redis.RedisConf         `yaml:"redis"`
}

// G Global configuration instance
var G *ConfYaml

//LoadResource
func (c *ConfYaml) LoadResource() {
	//Database resource loading
	mysql.Init(mysql.DefaultOrmType, c.MySQL)

	//Redis resource loading
	redis.Init(c.Redis)

	//Logger
	logger.Init(redis.Conn)

}

func fileExists(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// Parse parse yaml configuration file
func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !fileExists(cfg) {
		return fmt.Errorf("configuration file %s is nonexistent", cfg)
	}

	content, err := ioutil.ReadFile(cfg)
	if err != nil {
		return fmt.Errorf("read configuration file %s fail %s", cfg, err.Error())
	}

	var c ConfYaml
	if err = yaml.Unmarshal(content, &c); err != nil {
		return fmt.Errorf("parse configuration file %s fail %s", cfg, err.Error())
	}

	G = &c

	//
	G.LoadResource()

	return nil
}
