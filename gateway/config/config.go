/**
* @file: config.go ==> gateway/config
* @package: config
* @author: jingxiu
* @since: 2022/11/8
* @desc: 读取配置文件
 */

package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type (
	Config struct {
		Gateway   Gateway `yaml:"Gateway"`
		DB        DB      `yaml:"DB"`
		Rpc       Rpc     `yaml:"Rpc"`
		SecretKey string  `yaml:"SecretKey"`
		Jwt       JWT     `yaml:"Jwt"`
		Logger    Logger  `yaml:"Logger"`
	}
	// Gateway 网关启动地址
	Gateway struct {
		ServerName string `yaml:"ServerName"`
		Version    string `yaml:"Version"`
		Listen     string `yaml:"Listen"`
	}
	// DB  数据库链接配置
	DB struct {
		Type    string `yaml:"Type"`    // 链接类型
		Version string `yaml:"Version"` // 版本号
		Source  string `yaml:"Source"`  // 链接dns地址
	}
	// Cache 缓存链接配置
	Cache struct {
		Type    string `yaml:"Type"`    // 链接类型
		Version string `yaml:"Version"` // 版本号
		Source  string `yaml:"Source"`  // 链接dns地址
	}
	// Rpc 链接配置
	Rpc struct {
		// TODO register your gRpc Client config
	}
	JWT struct {
		Secret string `yaml:"Secret"`
		Expire string `yaml:"Expire"`
	}
	Logger struct {
		FilePath string `yaml:"FilePath"`
		FileName string `yaml:"FileName"`
	}
)

var C *Config

// ConfigInit 初始化配置项
func ConfigInit(path string) error {
	var err error
	C, err = ReadYamlConfig(path)
	return err
}

// ReadYamlConfig 读取配置文件
func ReadYamlConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		err = yaml.NewDecoder(f).Decode(conf)
		if err != nil {
			return nil, err
		}
	}
	return conf, nil
}
