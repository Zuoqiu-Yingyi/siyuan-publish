package server

import (
	"net/http"

	"publish/config"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.Redirect(http.StatusFound, config.C.Server.Index.URL)
}
