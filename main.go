package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	SetupSchema()

	router := gin.Default()
	router.POST("/register", Register)
	router.POST("/assign-dbcred", AssignCredentials)
	router.POST("/:list", ListRecords)

	router.Run("localhost:8080")
}
