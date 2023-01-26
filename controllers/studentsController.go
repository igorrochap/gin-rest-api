package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igorrochap/gin-rest-api/database"
	"github.com/igorrochap/gin-rest-api/models"
)

func AllStudents(ctx *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	ctx.JSON(200, students)
}

func Greeting(ctx *gin.Context) {
	name := ctx.Params.ByName("name")
	ctx.JSON(200, gin.H{
		"message": "Hello " + name + ", how you doing?",
	})
}

func CreateStudent(ctx *gin.Context) {
	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Request failed",
			"error":   err.Error(),
		})
	}
	database.DB.Create(&student)
	ctx.JSON(http.StatusCreated, student)
}

func GetStudentById(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, student)
}
