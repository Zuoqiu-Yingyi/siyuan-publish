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
