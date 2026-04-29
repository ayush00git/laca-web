package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ayush00git/laca-web/models"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourseHandler struct {
	Collection *mongo.Collection
}

func (h *CourseHandler) RegisterNewCourse (c *gin.Context) {
	var input models.Course

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// user registering the course should be an admin

	newCourse := models.Course{
		ID: primitive.NewObjectID(),
		Code: input.Code,
		MaxSeats: input.MaxSeats,
		// Students: , // decode the student ID from the jwt token
		CreatedAt: time.Now(),
	}

	_, err := h.Collection.InsertOne(context.Background(), newCourse)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			if strings.Contains(err.Error(), "code") {
				c.JSON(400, gin.H{"error": "This code is already assigned to an existing course"})
				return
			}
		}
		c.JSON(500, gin.H{"error": "Failed to register new course", "details": err.Error()})
		return
	}

	// for server logs
	fmt.Printf("registered a new course: %s", input.Code)

	c.JSON(201, gin.H{
		"message": "course updated successfully!",
		"course": input,
	})
}
