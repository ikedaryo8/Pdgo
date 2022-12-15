package main

import (
	//"net/http"

	"net/http"

	"github.com/gin-gonic/gin"
)

// type FormData struct {
// 	UserName    string `from : "UserName"`
// 	MailAddress string `from : "MailAddress" `
// 	Age         string `from : "Age"`
// 	Gender      string `from : "Gender"`
// 	OnBusstop   string `from : "OnBusstop"`
// 	OffBusstop  string `from : "OffBusstop"`
// }

func main() {
	router := gin.Default()
	router.Static("/static", "./static") //" URL", 格納場所
	router.LoadHTMLGlob("static/*")

	router.GET("/dev", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	//
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "static/index.html", gin.H{})
	})

	router.POST("/PostForm", func(c *gin.Context) {
		UserName := c.PostForm("User_name")
		UserMailAddress := c.PostForm("mail")
		UserAge := c.PostForm("age")
		UserGender := c.PostForm("gender")
		bus
		c.String(http.StatusOK, message)
	})
	router.Run(":8080")

}
