package main

import (
	"net/http"

	"macronomicon/api/controller"
	"macronomicon/api/database"

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

	println("START FETCHES")
	router.GET("/", index)
	println("END FETCH 1")
	router.GET("/macro_entries", controller.ListMacroEntries)
	println("END FETCH 2")
	router.GET("/users", controller.ListUsers)
	println("END FETCH 3")

	println("START SERVER")
	router.Run("localhost:8080")
	println("SERVING")
}
