package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Id    uint
	Name  string
	Email string
}

func dbConn() *gorm.DB {
	dsn := "root:787898@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connection established")
	return db
}

func getIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func postIndexHandler(c *gin.Context) {
	db := dbConn()
	name := c.PostForm("name")
	email := c.PostForm("email")
	user := User{
		Name:  name,
		Email: email,
	}
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "index.html", nil)
}

func userHandler(c *gin.Context) {
	var user []User
	db := dbConn()
	db.Find(&user)
	c.JSON(http.StatusOK, user)
}

func main() {
	db := dbConn()
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Created")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", getIndexHandler)
	router.POST("/", postIndexHandler)
	router.GET("/user", userHandler)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
