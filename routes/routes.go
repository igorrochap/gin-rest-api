package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igorrochap/gin-rest-api/controllers"
)

func HandleRequests() {
	server := gin.Default()
	server.GET("/students", controllers.AllStudents)
	server.GET("/:name", controllers.Greeting)
	server.POST("/students", controllers.CreateStudent)
	server.GET("/students/:id", controllers.GetStudentById)
	server.Run(":8000")
}
