package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/Get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "message from gin get method"})
	})

	r.POST("/Post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "From post method"})
	})

	r.Run(":6060")
}
