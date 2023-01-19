package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	fmt.Printf("Name: %s\nEmail: %s\n", name, email)
	c.HTML(200, "registration.html", nil)
}

func getFormHandler(c *gin.Context) {
	c.HTML(200, "registration.html", nil)
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", getFormHandler)
	router.POST("/", indexHandler)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
