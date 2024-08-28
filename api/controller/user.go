package controller

import (
	"net/http"

	"macronomicon/api/database"
	"macronomicon/api/model"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	var users []*model.User
	err := database.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
