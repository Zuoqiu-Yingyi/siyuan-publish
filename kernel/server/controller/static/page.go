package static

import (
	"publish/models"
	"publish/server/view"
	"publish/server/view/status"

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

	view.Block(c, doc)
}
