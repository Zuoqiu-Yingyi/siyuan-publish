package main

import (
	"flag"
	"fmt"
	"os"

	"publish/client"
	"publish/config"
	"publish/models"
	"publish/server"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	var (
		workspace string // 工作目录
		path      string // 配置文件路径
	)

	/* 获得当前工作目录 */
	if wd, err := os.Getwd(); err == nil {
		workspace = wd
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	/* 解析命令行参数 */
	flag.StringVar(&workspace, "workspace", workspace, `working directory`)
	flag.StringVar(&path, "config", "./default.config.toml", `config file path (*.config.toml")`)
	flag.Parse()
	// fmt.Println(path)

	/* 设置工作目录 */
	if err := os.Chdir(workspace); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/* 输出启动信息 */
	figure.NewFigure("siyuan  publish", "doom", true).Print()
	fmt.Printf("\n- workspace: %s\n- config: %s\n", workspace, path)

	/* 初始化配置文件 */
	config.Init(path)

	/* 初始化 HTTP 客户端 */
	client.Init()

	/* 初始化数据库 */
	models.Init()

	/* 初始化 Web 服务 */
	server.Init()

	/* 运行 Web 服务 */
	router := server.Server()
	router.Run(fmt.Sprintf(":%d", config.C.Server.Port))
}
