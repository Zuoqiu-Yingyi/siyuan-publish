package static

import (
	"net/http"
	"publish/config"
	"publish/models"
	"publish/server/status"

	"github.com/gin-gonic/gin"
)

/* 解析路由参数 id */
func Block(c *gin.Context) {
	root_id := c.GetString("root_id")
	doc := &models.Doc{}
	if d := doc.One(root_id); d == nil { // 数据库中没有该文档信息, 通过 API 查询
		status.S.StatusBlockNotFound(c)
		return
	}

	c.HTML(http.StatusOK, "block.html", gin.H{
		"Path":     doc.Path,
		"Hpath":    doc.Hpath,
		"Tag":      doc.Tag,
		"Icon":     doc.Icon,
		"Title":    doc.Title,
		"TitleImg": doc.TitleImg,
		"Content":  doc.Dom,
		"Render":   config.C.Render,
	})
}
