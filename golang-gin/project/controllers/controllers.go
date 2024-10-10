package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/userbarbosa/golang-alura/golang-gin/project/v2/models"
)

func GetAllStudents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Students)
}

func WelcomeMessage(ctx *gin.Context) {
	name := ctx.Params.ByName("name")

	ctx.JSON(http.StatusOK, gin.H{
		"API says": "Welcome " + name + ", how are you?",
	})
}
