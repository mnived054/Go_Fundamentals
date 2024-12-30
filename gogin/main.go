package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define POST route to accept User data
	r.POST("/create-user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User created", "user": user})

	})

	r.Run(":8089")
}
