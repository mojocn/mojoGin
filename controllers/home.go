package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {

	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}
