package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// mongodb connection

	// handlers and routes
	// r.Default

	// connection
	// log.Fatal()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "gotcha!",
		})
	})
	r.Run(":8080")
}
