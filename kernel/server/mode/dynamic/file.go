package dynamic

import (
	"net/http"
	"path"
	"strings"

	"publish/client"
	"publish/config"

	"github.com/gin-gonic/gin"
)

func File(c *gin.Context) {
	var (
		request_fullpath = c.FullPath()    // 资源请求路径
		relative_path    = c.Param("path") // 资源相对路径
		root_path        string            // 根路径
		file_path        string            // 文件路径
	)

	/* 设置资源存放路径 */
	switch {
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Appearance.Path):
		root_path = config.C.Server.Static.Appearance.Path
		file_path = path.Join(config.C.Server.Static.Appearance.FilePath, relative_path)
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Assets.Path):
		root_path = config.C.Server.Static.Assets.Path
		file_path = path.Join(config.C.Server.Static.Assets.FilePath, relative_path)
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Emojis.Path):
		root_path = config.C.Server.Static.Emojis.Path
		file_path = path.Join(config.C.Server.Static.Emojis.FilePath, relative_path)
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Export.Path):
		root_path = config.C.Server.Static.Export.Path
		file_path = path.Join(config.C.Server.Static.Export.FilePath, relative_path)
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Stage.Path):
		root_path = config.C.Server.Static.Stage.Path
		file_path = path.Join(config.C.Server.Static.Stage.FilePath, relative_path)
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Widgets.Path):
		root_path = config.C.Server.Static.Widgets.Path
		file_path = path.Join(config.C.Server.Static.Widgets.FilePath, relative_path)
	}

	// fmt.Println(request_fullpath, relative_path, file_path)

	/* 下载资源文件 */
	// REF [下载 - Req](https://req.cool/zh/docs/tutorial/download/)
	if _, err := client.C.R().SetOutputFile(file_path).Get(config.C.Siyuan.Server + root_path + relative_path); err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	/* 返回资源文件 */
	c.File(file_path)
}
