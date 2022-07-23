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

func (*BaseStatus) StatusInternalServerError(c *gin.Context, msg string) {
	code := http.StatusInternalServerError
	c.HTML(code, "error.html", gin.H{
		"StatusCode": code,
		"Message":    msg,
	})
}

func (*BaseStatus) StatusNotFound(c *gin.Context, msg string) {
	code := http.StatusNotFound
	c.HTML(code, "error.html", gin.H{
		"StatusCode": code,
		"Message":    msg,
	})
}

func (*BaseStatus) StatusBadRequest(c *gin.Context, msg string) {
	code := http.StatusBadRequest
	c.HTML(code, "error.html", gin.H{
		"StatusCode": code,
		"Message":    msg,
	})
}

func (*BaseStatus) StatusForbidden(c *gin.Context, msg string) {
	code := http.StatusForbidden
	c.HTML(code, "error.html", gin.H{
		"StatusCode": code,
		"Message":    msg,
	})
}

/* 思源内核服务错误 */
func (s *Status) StatusSiyuanServerError(c *gin.Context, msg string) {
	s.StatusInternalServerError(c, "Siyuan Serve Error")
}

/* 发布服务错误 */
func (s *Status) StatusPublishServerError(c *gin.Context) {
	s.StatusInternalServerError(c, "Publish Serve Error")
}

/* 未找到块 */
func (s *Status) StatusBlockNotFound(c *gin.Context) {
	s.StatusNotFound(c, "Block Not Found")
}

/* 参数错误 */
func (s *Status) StatusParamsError(c *gin.Context) {
	s.StatusBadRequest(c, "Parameter Error")
}

/* 拒绝访问 */
func (s *Status) StatusAccessDenied(c *gin.Context) {
	s.StatusForbidden(c, "Access Denied")
}
