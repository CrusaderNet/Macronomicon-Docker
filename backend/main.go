package main

import (
	"net/http"

	"macronomicon/backend/controller"
	"macronomicon/backend/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	err := database.Init()
	if err != nil {
		panic(err)
	}

	router.GET("/", index)
	router.GET("/macro_entries", controller.ListMacroEntries)
	router.GET("/users", controller.ListUsers)

	router.POST("/macro_entry", controller.InsertMacroEntry)

	router.Run("0.0.0.0:8080")
}
