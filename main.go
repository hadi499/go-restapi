package main

import (
	"go-rest-api/controllers"
	"go-rest-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()
	db := database.DB
	fileController := &controllers.FileController{DB: db}

	r.GET("/api/posts", controllers.Index)
	r.GET("/api/posts/:id", controllers.Detail)
	r.POST("/api/posts", fileController.Create)
	r.PUT("/api/posts/:id", fileController.Update)
	r.DELETE("/api/posts/:id", fileController.Delete)
	r.Run()
}
