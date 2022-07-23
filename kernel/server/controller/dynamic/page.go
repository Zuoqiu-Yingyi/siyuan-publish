package dynamic

import (
	"net/http"
	"strings"

	"publish/client"
	"publish/config"
	"publish/server/status"

	"github.com/gin-gonic/gin"
)

/* 解析路由参数 id */
func Block(c *gin.Context) {
	var (
		root_id string
		r       *client.ResponseBody
		err     error
		err_msg string
	)

	/* 获得指定文档的 DOM */
	root_id = c.GetString("root_id")
	r, err = client.GetBlockDomByID(client.C.R(), root_id, 0)
	r, err_msg = client.Response(r, err)
	if r == nil {
		status.S.StatusInternalServerError(c, err_msg)
		return
	}

	/* 渲染 HTML */
	blocks := r.Data.(map[string]interface{})["blocks"].([]interface{})
	switch {
	case len(blocks) == 0:
		status.S.StatusBlockNotFound(c)
		return
	default:
		block := blocks[0].(map[string]interface{})
		hpath := strings.Split(block["hPath"].(string), "/")
		title := hpath[len(hpath)-1]
		c.HTML(http.StatusOK, "block.html", gin.H{
			"Title":   title,
			"Content": block["content"],
			"Render":  config.C.Render,
		})
	}
}
