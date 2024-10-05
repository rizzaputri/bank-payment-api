package main

import (
	"github.com/gin-gonic/gin"
	"test-mnc/controllers"
	"test-mnc/initializers"
	"test-mnc/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/auth/register", controllers.SignUp)
	r.POST("/api/v1/auth/login", controllers.LogIn)
	r.POST("/api/v1/payments", middlewares.RequireAuth, controllers.CreatePayment)
	r.GET("/api/v1/histories", middlewares.RequireAuth, controllers.GetHistory)
	r.PUT("/api/v1/auth/logout", middlewares.RequireToken, controllers.LogOut)

	err := r.Run()
	if err != nil {
		panic("Failed to run the program")
	}
}
