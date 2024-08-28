package controller

import (
	"net/http"

	"macronomicon/api/database"
	"macronomicon/api/model"

	"github.com/gin-gonic/gin"
)

func ListMacroEntries(c *gin.Context) {
	var macroEntries []*model.MacroEntry
	err := database.DB.Find(&macroEntries).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, macroEntries)
}
