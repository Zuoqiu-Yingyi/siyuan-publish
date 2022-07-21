package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseStatus struct{}

type Status struct {
	BaseStatus
}

var (
	S *Status
)

func init() {
	S = &Status{}
}

func (*BaseStatus) StatusInternalServerError(c *gin.Context, info string) {
	code := http.StatusInternalServerError
	c.HTML(code, "error.html", gin.H{
		"StatusCode": code,
		"Title":      info,
	})
}

func (*BaseStatus) StatusNotFound(c *gin.Context, info string) {
	code := http.StatusNotFound
	c.HTML(code, "error.html", gin.H{
		"StatusCode": code,
		"Title":      info,
	})
}

func (*BaseStatus) StatusBadRequest(c *gin.Context, info string) {
	code := http.StatusBadRequest
	c.HTML(code, "error.html", gin.H{
		"StatusCode": code,
		"Title":      info,
	})
}

func (*BaseStatus) StatusForbidden(c *gin.Context, info string) {
	code := http.StatusForbidden
	c.HTML(code, "error.html", gin.H{
		"StatusCode": code,
		"Title":      info,
	})
}

func (s *Status) StatusSiyuanServerError(c *gin.Context) {
	s.StatusInternalServerError(c, "Siyuan Serve Error")
}

func (s *Status) StatusPublishServerError(c *gin.Context) {
	s.StatusInternalServerError(c, "Publish Serve Error")
}

func (s *Status) StatusBlockNotFound(c *gin.Context) {
	s.StatusNotFound(c, "Block Not Found")
}

func (s *Status) StatusParamsError(c *gin.Context) {
	s.StatusBadRequest(c, "Parameter Error")
}

func (s *Status) StatusAccessDenied(c *gin.Context) {
	s.StatusForbidden(c, "Access Denied")
}
