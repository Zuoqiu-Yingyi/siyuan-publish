package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
)

func Parse(path *string, version string) error {
	var (
		v         bool   // 是否显示版本信息
		workspace string // 工作目录
	)

	/* 获得当前工作目录 */
	if wd, err := os.Getwd(); err != nil {
		return err
	} else {
		workspace = wd
	}

	/* 解析命令行参数 */
	flag.BoolVar(&v, "v", false, `show version`)
	flag.StringVar(&workspace, "workspace", workspace, `working directory`)
	flag.StringVar(path, "config", "./default.config.toml", `config file path (*.config.toml")`)
	flag.Parse()

	/* 输出版本信息并退出 */
	if v {
		fmt.Printf("Version: %s\n", version)
		os.Exit(0)
	}

	/* 设置工作目录 */
	if err := os.Chdir(workspace); err != nil {
		return err
	}

	/* 输出启动信息 */
	figure.NewFigure("siyuan  publish", "doom", true).Print()
	fmt.Printf(`
- version: %s
- workspace: %s
- config: %s
`, version, workspace, *path)
	return nil
}
