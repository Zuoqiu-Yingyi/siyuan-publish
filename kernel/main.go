package main

import (
	"fmt"
	"os"

	"publish/client"
	"publish/command"
	"publish/config"
	"publish/models"
	"publish/server"
)

const (
	VERSION = "0.1.3" // 版本信息
)

func main() {
	var (
		path string // 配置文件路径
	)
	if err := command.Parse(&path, VERSION); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/* 初始化配置文件 */
	config.Init(path)
	fmt.Println("Loading configuration file successfully.")

	/* 初始化 HTTP 客户端 */
	client.Init()
	fmt.Println("Initrializing client successfully.")

	/* 初始化数据库 */
	models.Init()
	fmt.Println("Initrializing database successfully.")

	/* 初始化 Web 服务 */
	server.Init()
	fmt.Println("Initrializing server successfully.")

	/* 运行 Web 服务 */
	router := server.Server()
	router.Run(fmt.Sprintf(":%d", config.C.Server.Port))
}
