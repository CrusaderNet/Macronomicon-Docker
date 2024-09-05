package controller

import (
	"net/http"

	"fmt"
	"time"

	"macronomicon/api/database"
	"macronomicon/api/model"

	"github.com/gin-gonic/gin"
)

func ListMacroEntries(c *gin.Context) {
	println("ListMacroEntries CALL")
	var macroEntries []*model.MacroEntry
	if err := database.DB.Find(&macroEntries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, macroEntries)
}

func InsertMacroEntry(c *gin.Context) {
	println("Insert Macro Entry")

	var entry model.MacroEntry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		println(err)
		return
	}
	fmt.Printf("%+v", entry)

	entry.SubmitTime = time.Now()

	if err := database.DB.Create(&entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		println(err)
		return
	}

	c.Status(200)
}
