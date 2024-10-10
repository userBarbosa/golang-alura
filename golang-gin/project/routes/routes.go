package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/userbarbosa/golang-alura/golang-gin/project/v2/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/welcome/:name", controllers.WelcomeMessage)

	apiPort := os.Getenv("API_PORT")
	r.Run(":" + apiPort)
}
