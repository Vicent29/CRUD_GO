package Controllers

import (
	"CRUD-api/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTables ... Get all Tables
func GetTables(c *gin.Context) {
	var table []Models.Table
	err := Models.GetAllTables(&table)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, table)
	}
}

// CreateTable ... Create Table
func CreateTable(c *gin.Context) {
	var table Models.Table
	c.BindJSON(&table)
	err := Models.CreateTable(&table)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, table)
	}
}

// GetTableByID ... Get the Table by id
func GetTableByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var table Models.Table
	err := Models.GetTableByID(&table, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, table)
	}
}

// UpdateTable ... Update the Table information
func UpdateTable(c *gin.Context) {
	var table Models.Table
	id := c.Params.ByName("id")
	err := Models.GetTableByID(&table, id)
	if err != nil {
		c.JSON(http.StatusNotFound, table)
	}
	c.BindJSON(&table)
	err = Models.UpdateTable(&table, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, table)
	}
}

// DeleteTable ... Delete the Table
func DeleteTable(c *gin.Context) {
	var table Models.Table
	id := c.Params.ByName("id")
	err := Models.DeleteTable(&table, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + id: "is deleted"})
	}
}
