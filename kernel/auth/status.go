package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusSiyuanServerError(c *gin.Context) {
	// c.String(http.StatusInternalServerError, "Siyuan Server Error")
	c.HTML(http.StatusInternalServerError, "error.html", gin.H{
		"Title":      "Siyuan Server Error",
		"StatusCode": http.StatusInternalServerError,
	})
}

func StatusBlockNotFound(c *gin.Context) {
	// c.String(http.StatusNotFound, "Block Not Found")
	c.HTML(http.StatusNotFound, "error.html", gin.H{
		"Title":      "Block Not Found",
		"StatusCode": http.StatusNotFound,
	})
}

func StatusParamsError(c *gin.Context) {
	// c.String(http.StatusBadRequest, "Parameter Error")
	c.HTML(http.StatusBadRequest, "error.html", gin.H{
		"Title":      "Parameter Error",
		"StatusCode": http.StatusBadRequest,
	})
}

func StatusAccessDenied(c *gin.Context) {
	// c.String(http.StatusForbidden, "Access Denied")
	c.HTML(http.StatusForbidden, "error.html", gin.H{
		"Title":      "Access Denied",
		"StatusCode": http.StatusForbidden,
	})
}
