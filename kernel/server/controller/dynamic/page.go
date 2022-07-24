package dynamic

import (
	"publish/models"
	"publish/server/view"
	"publish/server/view/status"

	"github.com/gin-gonic/gin"
)

/* 解析路由参数 id */
func Block(c *gin.Context) {
	/* 获得指定文档的 DOM */
	root_id := c.GetString("root_id")
	if doc, err := models.GetDoc(root_id); err != nil {
		status.S.StatusInternalServerError(c, err.Error())
		return
	} else {
		view.Block(c, doc)
	}
}
