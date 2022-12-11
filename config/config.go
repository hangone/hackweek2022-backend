package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config Data // 配置文件结构体

type Data struct {
	Database Database `yaml:"database"` // 数据库配置
	Jwt      Jwt      `yaml:"jwt"`      // JWT 配置
}

type Database struct {
	Host     string `yaml:"host"`     // 数据库地址
	Port     int    `yaml:"port"`     // 数据库端口
	User     string `yaml:"user"`     // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	Dbname   string `yaml:"dbname"`   // 数据库名
	SslMode  string `yaml:"sslMode"`  // ssl模式
	TimeZone string `yaml:"TimeZone"` // 时区
}
type Jwt struct {
	SigningKey string `yaml:"signingKey"` // 签名密钥
}

func Init() {
	viper.SetConfigType("yaml")              // 设置配置文件格式为YAML
	viper.SetConfigFile("./Data/config.yml") // 指定配置文件路径
	err := viper.ReadInConfig()              // viper 解析配置文件
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = viper.Unmarshal(&Config) // viper 解析配置文件到结构体
	if err != nil {
		log.Fatalln(err.Error())
	}
}
