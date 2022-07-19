package server

import (
	"net/http"

	"publish/client"
	"publish/config"

	"github.com/gin-gonic/gin"
)

func block(c *gin.Context) {
	root_id := c.Param("id")
	r, err := client.GetBlockDomByID(client.C.R(), root_id)
	r = client.Response(c, r, err)
	if r == nil {
		return
	}
	blocks := r.Data.(map[string]interface{})["blocks"].([]interface{})
	switch {
	case len(blocks) == 0:
		c.String(http.StatusNotFound, "Block Not Found")
	case len(blocks) > 1:
		c.String(http.StatusInternalServerError, "More than one block found")
	default:
		block := blocks[0].(map[string]interface{})
		c.HTML(http.StatusOK, "block.html", gin.H{
			"Title":   block["hPath"],
			"Content": block["content"],
			"Render":  config.C.Render,
		})
	}
}
