package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* 解析 URL 参数 id */
func ID(c *gin.Context) {
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
