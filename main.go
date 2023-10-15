package main

import (
	"go-rest-api/controllers"
	"go-rest-api/database"
	"go-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()
	db := database.DB
	fileController := &controllers.FileController{DB: db}

	r.GET("/api/posts", middlewares.JWTMiddleware(), controllers.Index)
	r.GET("/api/posts/:id", controllers.Detail)
	r.POST("/api/posts", fileController.Create)
	r.PUT("/api/posts/:id", fileController.Update)
	r.DELETE("/api/posts/:id", fileController.Delete)
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)
	r.GET("/api/logout", controllers.Logout)

	r.Run()
}
