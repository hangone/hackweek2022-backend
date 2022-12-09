package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config Data

type Data struct {
	Database Database `yaml:"database"`
	Jwt      Jwt      `yaml:"jwt"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	SslMode  string `yaml:"sslMode"`
	TimeZone string `yaml:"TimeZone"`
}
type Jwt struct {
	SigningKey string `yaml:"signingKey"`
}

func Init() {
	//导入配置文件
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./Data/config.yml")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err.Error())
	}
	//将配置文件读到结构体中
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Println(err.Error())
	}
}
