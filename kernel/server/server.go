package server

import (
	"html/template"
	"strings"

	"publish/config"

	"github.com/gin-gonic/gin"
)

// REF [siyuan/serve.go at master · siyuan-note/siyuan](https://github.com/siyuan-note/siyuan/blob/master/kernel/server/serve.go)
func Server() (router *gin.Engine) {
	router = gin.Default()

	// REF [Gin框架Gin渲染 - RandySun - 博客园](https://www.cnblogs.com/randysun/p/15626537.html)
	router.SetFuncMap(template.FuncMap{
		"html": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	router.LoadHTMLGlob(config.C.Server.Templates) // 加载模板文件

	/* 静态文件服务 */
	router.Static(config.C.Server.Static.Appearance.Path, config.C.Server.Static.Appearance.FilePath)
	router.Static(config.C.Server.Static.Assets.Path, config.C.Server.Static.Assets.FilePath)
	router.Static(config.C.Server.Static.Emojis.Path, config.C.Server.Static.Emojis.FilePath)
	router.Static(config.C.Server.Static.Stage.Path, config.C.Server.Static.Stage.FilePath)
	router.Static(config.C.Server.Static.Widgets.Path, config.C.Server.Static.Widgets.FilePath)
	router.Static(config.C.Server.Static.JavaScript.Path, config.C.Server.Static.JavaScript.FilePath)
	router.Static(config.C.Server.Static.CSS.Path, config.C.Server.Static.CSS.FilePath)

	b := router.Group("/block")
	{
		/* 请求重定向 */
		redirect := func(c *gin.Context) {
			// REF [重定向 | Gin Web Framework](https://gin-gonic.com/zh-cn/docs/examples/redirects/)
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, "/block", "", 1)
			router.HandleContext(c)
		}

		b.GET("/appearance/*path", redirect)
		b.GET("/assets/*path", redirect)
		b.GET("/emojis/*path", redirect)
		b.GET("/widgets/*path", redirect)
		b.GET("/export/*path", redirect)

		// REF [Query 和 post form | Gin Web Framework](https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/)
		b.GET("/", id)

		// REF [绑定 Uri | Gin Web Framework](https://gin-gonic.com/zh-cn/docs/examples/bind-uri/)
		b.GET("/:id", block)
	}
	return
}
