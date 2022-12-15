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
	router.LoadHTMLGlob("./static/*.tmpl")

	router.GET("/dev", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	//
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.tmpl", gin.H{})
	})
	router.GET("/reservation", func(ctx *gin.Context) {
		ctx.HTML(200, "form.tmpl", gin.H{})

	})

	// router.POST("/send", func(c *gin.Context) {
	// 	UserName := c.PostForm("User_name")
	// 	UserMailAddress := c.PostForm("mail")
	// 	UserAge := c.PostForm("age")
	// 	UserGender := c.PostForm("gender")
	// 	BusstopON := c.PostForm("board_bus_stop")
	// 	BusstopOFf := c.PostForm("get_off_bus_stop")
	// 	c.String(http.StatusOK)

	// })
	router.Run(":8080")

}
