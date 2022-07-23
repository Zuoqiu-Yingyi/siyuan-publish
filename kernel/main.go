package main

import (
	"flag"
	"fmt"

	"publish/client"
	"publish/config"
	"publish/server"
)

func main() {
	var (
		path string
	)

	/* 解析命令行参数 */
	flag.StringVar(&path, "config", "./default.config.toml", `config file path (*.config.toml")`)
	flag.Parse()
	// fmt.Println(path)

	/* 初始化配置文件 */
	config.Init(path)

	/* 初始化 HTTP 客户端 */
	client.Init()

	/* 初始化数据库 */
	models.Init()

	/* 初始化 Web 服务 */
	server.Init()
	router := server.Server()

	/* 运行 Web 服务 */
	router.Run(fmt.Sprintf(":%d", config.C.Server.Port))
}
