package server

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"publish/config"
	"publish/locale"
	"publish/server/controller"
	"publish/server/controller/cache"
	"publish/server/controller/dynamic"
	"publish/server/controller/static"
	"publish/server/middleware/auth/access"
	"publish/server/middleware/parser"

	"github.com/gin-gonic/gin"
)

/* 初始化 Web 服务配置 */
func Init() {
	/* 设置服务运行模式 */
	if config.C.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	/* 是否重置静态目录 */
	if config.C.Server.Static.Reset {
		if err := os.RemoveAll(config.C.Server.Static.Path); err != nil {
			panic(err)
		} else {
			fmt.Println(locale.T("static.success"))
		}
	}

	/* 同时将日志写入日志文件与控制台 */
	// REF [如何记录日志 | Gin Web Framework](https://gin-gonic.com/zh-cn/docs/examples/write-log/)
	now := time.Now()
	log_file_name := now.Format("2006-01-02") + ".log"
	log_file_path := path.Join(config.C.Server.Logs, log_file_name)

	/* 创建日志目录 */
	if err := os.MkdirAll(config.C.Server.Logs, config.DEFAULT_DIR_MODE); err != nil {
		panic(err)
	}

	// REF [打开或新建一个文件](https://blog.csdn.net/weixin_45193103/article/details/123479196)
	// REF [os - Constants](https://studygolang.com/static/pkgdoc/pkg/os.htm#pkg-index)
	if log_file, err := os.OpenFile(log_file_path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		panic(err)
	} else {
		gin.DefaultWriter = io.MultiWriter(log_file, os.Stdout)
	}
}

// REF [siyuan/serve.go at master · siyuan-note/siyuan](https://github.com/siyuan-note/siyuan/blob/master/kernel/server/serve.go)
func Server() (router *gin.Engine) {
	/* 引擎 */
	router = gin.Default()

	/* 注册自定义模板函数 */
	// REF [Gin框架Gin渲染 - RandySun - 博客园](https://www.cnblogs.com/randysun/p/15626537.html)
	router.SetFuncMap(template.FuncMap{
		/* 渲染 HTML 原文(不转义) */
		"html": func(str string) template.HTML {
			return template.HTML(str)
		},
		/* 判断字符串是否为 16 进制字符串 */
		"hex": func(str string) bool {
			// REF
			if _, err := strconv.ParseUint(str, 16, 64); err != nil {
				return false
			} else {
				return true
			}
		},
		/* 取模 */
		"mod": func(a, mod int) int {
			return a % mod
		},
		/* i18n */
		"Tr": locale.L.Tr,
	})

	/* 加载模板文件 */
	router.LoadHTMLGlob(config.C.Server.Templates)

	/* 主页 */
	for _, path := range config.C.Server.Index.Paths {
		router.GET(path, index)
	}

	/* 静态文件服务 */
	router.Static(config.C.Server.Static.JavaScript.Path, config.C.Server.Static.JavaScript.FilePath)
	router.Static(config.C.Server.Static.CSS.Path, config.C.Server.Static.CSS.FilePath)
	router.StaticFile(config.C.Server.Static.Favicon.Path, config.C.Server.Static.Favicon.FilePath)

	/* 资源文件加载模式 */
	switch config.C.Server.Mode.File {
	case "dynamic": // 动态加载
		router.GET(config.C.Server.Static.Appearance.Path+"/*path", dynamic.File)
		router.GET(config.C.Server.Static.Assets.Path+"/*path", dynamic.File)
		router.GET(config.C.Server.Static.Emojis.Path+"/*path", dynamic.File)
		router.GET(config.C.Server.Static.Snippets.Path+"/*path", dynamic.File)
		router.GET(config.C.Server.Static.Widgets.Path+"/*path", dynamic.File)
		router.GET(config.C.Server.Static.Export.Path+"/*path", dynamic.File)
		router.GET(config.C.Server.Static.Stage.Path+"/*path", dynamic.File)
	case "cache": // 动态缓存
		router.GET(config.C.Server.Static.Appearance.Path+"/*path", cache.File)
		router.GET(config.C.Server.Static.Assets.Path+"/*path", cache.File)
		router.GET(config.C.Server.Static.Emojis.Path+"/*path", cache.File)
		router.GET(config.C.Server.Static.Snippets.Path+"/*path", cache.File)
		router.GET(config.C.Server.Static.Widgets.Path+"/*path", cache.File)
		router.GET(config.C.Server.Static.Export.Path+"/*path", cache.File)
		router.GET(config.C.Server.Static.Stage.Path+"/*path", cache.File)
	case "static": // 静态加载
		fallthrough
	default:
		router.Static(config.C.Server.Static.Appearance.Path, config.C.Server.Static.Appearance.FilePath)
		router.Static(config.C.Server.Static.Assets.Path, config.C.Server.Static.Assets.FilePath)
		router.Static(config.C.Server.Static.Emojis.Path, config.C.Server.Static.Emojis.FilePath)
		router.Static(config.C.Server.Static.Snippets.Path, config.C.Server.Static.Snippets.FilePath)
		router.Static(config.C.Server.Static.Widgets.Path, config.C.Server.Static.Widgets.FilePath)
		router.Static(config.C.Server.Static.Export.Path, config.C.Server.Static.Export.FilePath)
		router.Static(config.C.Server.Static.Stage.Path, config.C.Server.Static.Stage.FilePath)
	}

	router_block := router.Group(config.C.Server.Pathname)
	{
		// /* 请求重定向 */
		// redirect := func(c *gin.Context) {
		// 	// REF [重定向 | Gin Web Framework](https://gin-gonic.com/zh-cn/docs/examples/redirects/)
		// 	c.Request.URL.Path = strings.Replace(c.Request.URL.Path, "/block", "", 1)
		// 	router.HandleContext(c)
		// }

		// /* 资源文件请求重定向 */
		// router_block.GET("/appearance/*path", redirect)
		// router_block.GET("/assets/*path", redirect)
		// router_block.GET("/emojis/*path", redirect)
		// router_block.GET("/widgets/*path", redirect)
		// router_block.GET("/export/*path", redirect)
		// router_block.GET("/stage/*path", redirect)

		/* 文档页面加载方式 */
		// 使用 URL 参数 id 跳转到指定的块
		// REF [Query 和 post form | Gin Web Framework](https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/)
		router_block.GET(
			"/",
			parser.QueryID,
			access.Access,
			controller.ID,
		)
		switch config.C.Server.Mode.Page {
		case "dynamic": // 动态加载
			// 请求指定的文档
			// REF [绑定 Uri | Gin Web Framework](https://gin-gonic.com/zh-cn/docs/examples/bind-uri/)
			router_block.GET(
				"/:id",
				parser.ParamID,
				access.Access,
				dynamic.Block,
			)
		case "cache": // 动态缓存
			router_block.GET(
				"/:id",
				parser.ParamID,
				access.Access,
				cache.Block,
			)
		case "static": // 静态加载
			fallthrough
		default:
			router_block.GET(
				"/:id",
				parser.ParamID,
				access.Access,
				static.Block,
			)
		}
	}
	return
}
