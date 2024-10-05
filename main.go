package main

import (
	"github.com/gin-gonic/gin"
	"test-mnc/controllers"
	"test-mnc/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/auth/register", controllers.SignUp)
	r.POST("/api/v1/auth/login", controllers.Login)

	err := r.Run()
	if err != nil {
		panic("Failed to run the program")
	}
}
