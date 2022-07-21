package server

import (
	"net/http"

	"publish/client"
	"publish/server/status"

	"github.com/gin-gonic/gin"
)

func id(c *gin.Context) {
	id := c.Query("id")
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
