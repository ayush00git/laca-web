package routes

import (
	"github.com/ayush00git/laca-web/handlers"
	"github.com/gin-gonic/gin"	
)

func CourseRoute (route *gin.Engine, courseHandler *handlers.CourseHandler) {
	api := route.Group("/api/admin")
	{
		api.POST("/register-course", courseHandler.RegisterNewCourse)
	}
}
