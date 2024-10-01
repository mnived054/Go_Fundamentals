package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/getData", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Hi i am GIN framework"})
	})

	router.POST("/getDataPOST", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "hi this POST message from GIN framework"})
	})

	router.Run(":8086")
}
