package dynamic

import (
	"net/http"

	"publish/client"
	"publish/config"
	"publish/server/status"

	"github.com/gin-gonic/gin"
)

type Page struct{}

var (
	P *Page
)

func init() {
	P = &Page{}
}

/* 解析 URL 参数 id */
func (*Page) ID(c *gin.Context) {
	id := c.GetString("id")
	r, err := client.GetBlockByID(client.C.R(), id)
	r = client.Response(c, r, err)
	if r == nil {
		status.S.StatusSiyuanServerError(c)
		return
	}
	data := r.Data.([]interface{})
	switch {
	case len(data) == 0:
		status.S.StatusBlockNotFound(c)
		return
	default:
		block := data[0].(map[string]interface{})
		switch block["type"].(string) {
		case "d":
			c.Request.URL.Path = "/block/" + id
			query := c.Request.URL.Query()
			query.Del("id")
			// fmt.Printf("%+v\n", query)
			// fmt.Printf("%+v\n", query.Encode())
			c.Request.URL.RawQuery = query.Encode()
		default:
			c.Request.URL.Path = "/block/" + block["root_id"].(string)
			// router.HandleContext(c)
		}
		c.Redirect(http.StatusMovedPermanently, c.Request.URL.String())
	}
}

/* 解析路由参数 id */
func (*Page) Block(c *gin.Context) {
	root_id := c.GetString("id")
	r, err := client.GetBlockDomByID(client.C.R(), root_id, 0)
	r = client.Response(c, r, err)
	if r == nil {
		status.S.StatusSiyuanServerError(c)
		return
	}
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
