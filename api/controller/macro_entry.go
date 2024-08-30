package controller

import (
	"net/http"

	"macronomicon/api/database"
	"macronomicon/api/model"

	"github.com/gin-gonic/gin"
)

func ListMacroEntries(c *gin.Context) {
	println("ListMacroEntries CALL")
	var macroEntries []*model.MacroEntry
	err := database.DB.Find(&macroEntries).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, macroEntries)
}
