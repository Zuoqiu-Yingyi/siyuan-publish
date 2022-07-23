package controller

import "github.com/gin-gonic/gin"

type Page interface {
	ID(c *gin.Context)
	Block(c *gin.Context)
}
