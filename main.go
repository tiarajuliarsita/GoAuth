package main

import (
	"lear-jwt/controller"
	"lear-jwt/initializers"
	"lear-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)
	r.GET("/logout", controller.LogOut)

	v1 := r.Group("/api")
	v1.Use(middleware.JwtMiddleWare()) // Memasukkan middleware sebelum rute-rute di dalam grup
	{
		v1.GET("/products", controller.Store)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
