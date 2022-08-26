package main

import (
	"fmt"
	"os"

	"publish/client"
	"publish/command"
	"publish/config"
	"publish/locale"
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
	locale.Init(config.C.I18n.Directory, config.C.I18n.Default)
	fmt.Println(locale.T("config.success"))

	/* 初始化 HTTP 客户端 */
	client.Init()
	fmt.Println(locale.T("client.success"))

	/* 初始化数据库 */
	models.Init()
	fmt.Println(locale.T("models.success"))

	/* 初始化 Web 服务 */
	server.Init()
	fmt.Println(locale.T("server.success"))

	/* 运行 Web 服务 */
	router := server.Server()
	router.Run(fmt.Sprintf(":%d", config.C.Server.Port))
}
