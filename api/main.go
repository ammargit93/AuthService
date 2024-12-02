package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupSchema()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	router.POST("/register", Register)
	router.POST("/assign-dbcred", AssignCredentials)
	router.POST("/login", Login)
	router.POST("/:list", ListRecords)
	router.GET("/protected", AuthMiddleware(), Protected)

	router.Run("localhost:8080")
}
