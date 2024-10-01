package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	c.JSON(200, gin.H{"data": "Hi i am GIN Framework",
		"bodyData": string(value),
	})
}

func getDataPost(c *gin.Context) {
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	c.JSON(200, gin.H{"data": "Hi i am POST method from  GIN Framework",
		"bodyData": string(value)})
}

func main() {
	router := gin.Default()

	router.GET("/getData", getData)
	router.POST("/getDataPost", getDataPost)

	router.Run(":8086")

}
