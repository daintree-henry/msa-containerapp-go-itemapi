package main

import (
	"github.com/daintree-henry/msa-containerapp-go-itemapi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/health", controllers.Health)
	router.GET("/api/items", controllers.ListItem)
	router.GET("/api/items/:id", controllers.GetItem)
	router.POST("/api/items", controllers.CreateItem)
	router.DELETE("/api/items/:id", controllers.DeleteItem)

	if err := router.Run(":80"); err != nil {
		panic(err)
	}
}
