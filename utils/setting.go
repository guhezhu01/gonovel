package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	AppMode    string
	HttpPort   string
	JwtKey     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

type Db struct {
	DbHost     string `yaml:"DbHost"`
	DbPort     string `yaml:"DbPort"`
	DbUser     string `yaml:"DbUser"`
	DbPassWord string `yaml:"DbPassWord"`
	DbName     string `yaml:"DbName"`
}

type Env struct {
	AppMode  string `yaml:"AppMode"`
	HttpPort string `yaml:"HttpPort"`
	JwtKey   string `yaml:"JwtKey"`
	Mysql    Db     `yaml:"Mysql"`
}

func init() {
	content, err := os.ReadFile("config/config.yaml")
	var env Env
	err = yaml.Unmarshal(content, &env)
	fmt.Println(env)
	if err != nil {
		fmt.Println("配置文件读取错误！")
	}
	LoadServer(env)
	LoadData(env)

}

func LoadServer(env Env) {
	AppMode = env.AppMode
	HttpPort = env.HttpPort
	JwtKey = env.JwtKey
}

func LoadData(env Env) {
	DbHost = env.Mysql.DbHost
	DbPort = env.Mysql.DbPort
	DbUser = env.Mysql.DbUser
	DbPassWord = env.Mysql.DbPassWord
	DbName = env.Mysql.DbName
}
