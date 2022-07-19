package server

import (
	"net/http"

	"publish/client"

	"github.com/gin-gonic/gin"
)

func id(c *gin.Context) {
	id := c.Query("id")
	r, err := client.GetBlockByID(client.C.R(), id)
	r = client.Response(c, r, err)
	if r == nil {
		return
	}
	data := r.Data.([]interface{})
	switch {
	case len(data) == 0:
		c.String(http.StatusNotFound, "Block Not Found")
	case len(data) > 1:
		c.String(http.StatusInternalServerError, "More than one block found")
	default:
		block := data[0].(map[string]interface{})
		switch block["type"].(string) {
		case "d":
			c.Redirect(http.StatusMovedPermanently, "/block/"+id)
		default:
			c.Request.URL.Path = "/block/" + block["root_id"].(string)
			c.Redirect(http.StatusMovedPermanently, c.Request.URL.String())
			// router.HandleContext(c)
		}
	}

}
