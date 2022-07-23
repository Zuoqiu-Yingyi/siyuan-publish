package dynamic

import (
	"net/http"

	"publish/client"
	"publish/config"
	"publish/server/controller"
	"publish/server/status"

	"github.com/gin-gonic/gin"
)

type page struct {
	controller.Page
}

var (
	P *page
)

func init() {
	P = &page{}
}

/* 解析 URL 参数 id */
func (*page) ID(c *gin.Context) {
	id := c.GetString("id")
	root_id := c.GetString("root_id")

	c.Request.URL.Path = "/block/" + root_id
	if id == root_id { // 查询文档块
		query := c.Request.URL.Query()
		query.Del("id")
		// fmt.Printf("%+v\n", query)
		// fmt.Printf("%+v\n", query.Encode())
		c.Request.URL.RawQuery = query.Encode()
		// router.HandleContext(c)
	}

	c.Redirect(http.StatusMovedPermanently, c.Request.URL.String())
}

/* 解析路由参数 id */
func (*page) Block(c *gin.Context) {
	var (
		root_id string
		r       *client.ResponseBody
		err     error
		err_msg string
	)

	/* 获得指定文档的 DOM */
	root_id = c.GetString("id")
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
		c.HTML(http.StatusOK, "block.html", gin.H{
			"Title":   block["hPath"],
			"Content": block["content"],
			"Render":  config.C.Render,
		})
	}
}
