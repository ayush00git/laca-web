package main

import (
	"fmt"

	"github.com/ayush00git/laca-web/db"
	"github.com/ayush00git/laca-web/handlers"
	"github.com/ayush00git/laca-web/helpers"
	"github.com/ayush00git/laca-web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// mongodb connection
	uri := helpers.GetEnvVar("MONGO_URI")

	database, err := db.ConnectToMongo(uri)
	if err != nil {
		fmt.Printf("Connection to mongodb is failing: %s", err)
	}

	courseCollection := database.Collection("courses")
	courseHandler := &handlers.CourseHandler{
		Collection: courseCollection,
	}

	r := gin.Default()

	routes.CourseRoute(r, courseHandler)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "gotcha!",
		})
	})
	fmt.Println("Server running at port 8080")
	r.Run(":8080")
}
