package main

import (
	"main/controllers"
	"main/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleWare())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
