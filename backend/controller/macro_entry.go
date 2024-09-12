package controller

import (
	"net/http"

	"fmt"

	"macronomicon/backend/database"
	"macronomicon/backend/model"

	"github.com/gin-gonic/gin"
)

func ListMacroEntries(c *gin.Context) {
	println("ListMacroEntries CALL")
	var macroEntries []*model.MacroEntry
	if err := database.DB.Order("submit_time desc").Find(&macroEntries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, macroEntries)
}

func InsertMacroEntry(c *gin.Context) {
	println("Insert Macro Entry")

	var req model.InsertMacroEntryRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		println(err)
		return
	}

	entry := model.MacroEntry{
		UserID:   req.UserID,
		Proteins: req.Proteins,
		Carbs:    req.Carbs,
		Fats:     req.Fats,
	}

	fmt.Printf("%+v", entry)

	if err := database.DB.Create(&entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		println(err)
		return
	}

	c.Status(200)
}
