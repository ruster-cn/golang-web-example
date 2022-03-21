package pkg

import (
	"fmt"
	"io/ioutil"

	"github.com/op-server/pkg/server/config"
	"github.com/op-server/pkg/server/orm"
	log "github.com/ruster-cn/zap-log-wrapper"
	"gopkg.in/yaml.v2"
)

type ConfigurationInterface interface {
	Default() ConfigurationInterface
	Validate() error
}

type Configuration struct {
	PaasServer *config.HTTPServerConfiguration `yaml:"paas_server"`
	Log        *log.LoggerConfiguration        `yaml:"log"`
	Orm        *orm.MysqlConnectConfiguration  `yaml:"orm"`
}

//TODO: 每个package定义自己的config,在这里组合到一起
func NewConfigurationFromFile(file string) (*Configuration, error) {
	config, err := loadConfig(file)
	if err != nil {
		return nil, fmt.Errorf("new config fail,%v", err)
	}
	if err := config.PaasServer.Default().Validate(); err != nil {
		return nil, fmt.Errorf("new paas server config fail,%v", err)
	}
	if err := config.Log.Validate(); err != nil {
		return nil, fmt.Errorf("new log config fail,%v", err)
	}
	return config, nil
}

func loadConfig(file string) (*Configuration, error) {
	var c Configuration
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read config file fail,%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("read config file fail,%v", err)
	}
	return &c, nil
}
