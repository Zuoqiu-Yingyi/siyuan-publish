package client

import (
	"time"

	"publish/config"

	"github.com/imroc/req/v3"
)

var C *req.Client

func init() {
}

func InitClient() {
	C = req.C().
		// DevMode().
		// SetCommonContentType("application/json").
		SetCommonHeader("Authorization", "Token "+config.C.Siyuan.Token).
		SetTimeout(time.Duration(config.C.Siyuan.Timeout) * time.Millisecond)
}
