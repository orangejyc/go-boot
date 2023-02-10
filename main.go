package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const CFG_KEY_GO_BOOT string = "goboot.env"

const DEFAULT_VALUE_GO_BOOT_ENV string = "default"

const DEFAULT_VALUE_GO_BOOT_APP_NAME string = "goboot"

const DEFAULT_VALUE_CURRENT_DIR string = "."

const DEFAULT_VALUE_CFG_FILE_SPILT string = "-"

const DEFAULT_VALUE_CFG_TYPE string = "yaml"

func main() {

	// TODO 收取环境变量&命令行参数 goboot.env

	// 默认环境ddefault
	app_env := DEFAULT_VALUE_GO_BOOT_ENV
	// 默认应用名称go-boot
	app_name := DEFAULT_VALUE_GO_BOOT_APP_NAME
	//默认配置文件读取路径
	cfg_path := DEFAULT_VALUE_CURRENT_DIR
	cfg_file_name := app_name
	if app_env != DEFAULT_VALUE_GO_BOOT_ENV {
		cfg_file_name = app_name + DEFAULT_VALUE_CFG_FILE_SPILT + app_env
	}

	viper.SetConfigName(cfg_file_name)
	viper.AddConfigPath(cfg_path)
	viper.SetConfigType(DEFAULT_VALUE_CFG_TYPE)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		// TODO 需要指定托底日志打印到当前目录,或者直接异常抛出
		log.Fatal(err)
	}

	// TODO 判断当前环境

	// TODO viper根据环境参数读取配置文件

	// TODO 覆盖配置文件配置项

	// TODO 根据配置初始化日志配置

	// TODO 初始化持久连接

	// TODO 初始化di容器,如果是web,需要再启动web

	// TODO 初始化路由

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		/*
			c.JSON(200, gin.H{
				"message": "pong",
			})*/

		panic("gin-test")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
