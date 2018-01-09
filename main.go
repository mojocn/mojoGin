package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/mojoGin/controllers"
	"github.com/mojocn/mojoGin/models"
)

func main() {
	//init Redis
	defer models.DB.Close()

	router := gin.Default()
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.

	router.GET("/", controllers.HomeGet)

	router.Run(":8080")
}
