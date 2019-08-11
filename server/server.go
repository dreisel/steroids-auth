package server

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "v1.0.0",
		})
	})
	return r
}
