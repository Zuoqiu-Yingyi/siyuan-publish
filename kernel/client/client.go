package client

import (
	"time"

	"publish/config"
	"publish/utils"

	"github.com/imroc/req/v3"
)

var C *req.Client

func Init() {
	C = req.C().
		// SetCommonContentType("application/json").
		SetCommonHeader("Authorization", "Token "+config.C.Siyuan.Token).
		SetTimeout(time.Duration(config.C.Siyuan.Timeout) * time.Millisecond).
		SetCommonRetryCount(config.C.Siyuan.Retry).
		SetBaseURL(config.C.Siyuan.Server)

	if config.C.Siyuan.Debug { // 是否启动客户端调试模式
		C.DevMode()
	}

	handle(config.C)
}

/* 处理配置文件 */
func handle(conf *config.Config) {
	/* 基础样式文件自动解析 */
	if conf.Render.File.Style.Base == "" {
		/* 从 index.html 中解析 base.hash.css 文件名 */
		if index, err := Get(C.R(), config.C.Server.Static.Stage.Path+"/build/mobile/index.html"); err == nil {
			file_name := utils.GetStyleFileName(index.String())
			conf.Render.File.Style.Base = config.C.Server.Static.Stage.Path + "/build/mobile/" + file_name
		} else {
			panic(err)
		}
	}
}
