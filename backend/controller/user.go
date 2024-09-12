package controller

import (
	"net/http"

	"macronomicon/backend/database"
	"macronomicon/backend/model"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	println("ListUsers CALL")
	var users []*model.User
	err := database.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
