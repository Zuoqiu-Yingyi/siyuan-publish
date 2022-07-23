package access

import (
	"publish/config"

	"github.com/gin-gonic/gin"
)

/*
判断文档是否可访问
REF [siyuan/session.go at master · siyuan-note/siyuan](https://github.com/siyuan-note/siyuan/blob/master/kernel/model/session.go)
*/
func Access(c *gin.Context) {
	var (
		access   bool                 // 是否可访问
		f_before func(c *gin.Context) // 回调函数
		f_after  func(c *gin.Context) // 回调函数
	)

	switch config.C.Server.Mode.Page {
	case "dynamic":
		access, f_before, f_after = Dynamic(c)
	case "cache":
		access, f_before, f_after = Cache(c)
	case "static":
		access, f_before, f_after = Static(c)
	}

	if f_before != nil {
		f_before(c)
	}

	if access {
		c.Next()
	} else {
		c.Abort()
	}

	if f_after != nil {
		f_after(c)
	}
}
