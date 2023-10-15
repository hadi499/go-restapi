package main

import (
	"go-rest-api/controllers"
	"go-rest-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	r.GET("/api/posts", controllers.Index)
	r.GET("/api/posts/:id", controllers.Detail)
	r.POST("/api/posts", controllers.Create)
	r.PUT("/api/posts/:id", controllers.Update)
	r.DELETE("/api/posts", controllers.Delete)
	r.Run()
}
