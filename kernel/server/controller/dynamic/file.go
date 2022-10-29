package dynamic

import (
	"bytes"
	"net/http"
	"strings"

	"publish/client"
	"publish/config"

	"github.com/gin-gonic/gin"
)

/* 资源文件转发 */
func File(c *gin.Context) {
	var (
		request_fullpath = c.FullPath()    // 资源请求路径
		relative_path    = c.Param("path") // 资源相对路径
		//  REF [golang中的io.Reader/Writer_SuPhoebe的博客-CSDN博客_golang io.reader](https://blog.csdn.net/u013007900/article/details/89126811
		buffer    = new(bytes.Buffer) // 文件缓冲区
		root_path string              // 根路径
	)

	/* 设置资源存放路径 */
	switch {
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Appearance.Path):
		root_path = config.C.Server.Static.Appearance.Path
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Assets.Path):
		root_path = config.C.Server.Static.Assets.Path
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Emojis.Path):
		root_path = config.C.Server.Static.Emojis.Path
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Export.Path):
		root_path = config.C.Server.Static.Export.Path
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Snippets.Path):
		root_path = config.C.Server.Static.Snippets.Path
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Stage.Path):
		root_path = config.C.Server.Static.Stage.Path
	case strings.HasPrefix(request_fullpath, config.C.Server.Static.Widgets.Path):
		root_path = config.C.Server.Static.Widgets.Path
	}

	/* 获取资源文件 */
	// REF [下载 - Req](https://req.cool/zh/docs/tutorial/download/)
	if response, err := client.C.R().SetOutput(buffer).Get(root_path + relative_path); err != nil || response.IsError() {
		c.Status(http.StatusNotFound)
	} else {
		// [从 reader 读取数据 | Gin Web Framework](https://gin-gonic.com/zh-cn/docs/examples/serving-data-from-reader/)
		// c.Data(http.StatusOK, response.GetContentType(), buffer.Bytes())
		c.DataFromReader(http.StatusOK, response.ContentLength, response.GetContentType(), buffer, nil)
	}

	/* 返回资源文件 */
	// REF [gin框架中多种数据格式返回请求结果 - 码农教程](http://www.manongjc.com/detail/26-epvqgdbxefbfjfd.html)
	// if _, err := c.Writer.Write(buffer.Bytes()); err != nil {
	// 	status.S.StatusPublishServerError(c)
	// }

}
