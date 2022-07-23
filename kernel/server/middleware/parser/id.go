package parser

import "github.com/gin-gonic/gin"

/* 从 URL 请求参数中获得 ID */
func QueryID(c *gin.Context) {
	c.Set("id", c.Query("id"))
	c.Next()
}

/* 从 URL 路径参数中获得 ID */
func ParamID(c *gin.Context) {
	c.Set("id", c.Param("id"))
	c.Next()
}
