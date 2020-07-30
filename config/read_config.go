package config

import (
	"godotenv"
)

// 读取配置文件并加载到内存中
func init() {
	ReadLocalConfig()
	ReadExternalConfig()
}

// 本地文件读取
// load 读取当前文件夹所有.env文件并加载env环境变量
func ReadLocalConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

// 远程配置中心读取
func ReadExternalConfig() {}
