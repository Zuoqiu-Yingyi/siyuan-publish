package client

import (
	"time"

	"publish/config"

	"github.com/imroc/req/v3"
)

var C *req.Client

func init() {
}

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
}
