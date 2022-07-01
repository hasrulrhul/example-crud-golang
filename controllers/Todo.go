package controllers

import (
	"example-crud-golang/config"
	"example-crud-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var todo []models.Todo
	config.DB.Find(&todo)
	c.JSON(http.StatusOK, todo)
}

func Create(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		panic(err)
	}
	config.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	config.DB.First(&todo, id)
	c.JSON(http.StatusOK, todo)
}

func Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := config.DB.First(&todo, id).Error
	if err != nil {
		panic(err)
	}

	if err := c.BindJSON(&todo); err != nil {
		panic(err)
	}

	config.DB.Updates(&todo)
	c.JSON(http.StatusOK, todo)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := config.DB.First(&todo, id).Error
	if err != nil {
		panic(err)
	}
	config.DB.Delete(&todo)
	c.JSON(http.StatusOK, todo)
}
