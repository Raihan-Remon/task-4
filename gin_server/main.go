package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func indexHandler(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}
func aboutHandler(ctx *gin.Context) {
	ctx.HTML(200, "about.html", nil)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", indexHandler)
	router.GET("/about", aboutHandler)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
