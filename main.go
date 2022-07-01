package main

import (
	"example-crud-golang/config"
	"example-crud-golang/controllers"
	"example-crud-golang/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Todo{})
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/todo", controllers.Index)
	router.POST("/todo", controllers.Create)
	router.GET("/todo/:id", controllers.Show)
	router.PUT("/todo/:id", controllers.Update)
	router.DELETE("/todo/:id", controllers.Delete)

	router.Run()
}
