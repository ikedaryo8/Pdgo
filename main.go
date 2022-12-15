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
	// 起動確認用
	router.GET("/dev", test)
	// トップページを表示
	router.GET("/", getTopPage)
	// 入力用フォームを表示
	router.GET("/reservation", getForm)
	// フォーム送信
	router.POST("/reservation", postForm)
	// router.POST("/sendform", func(c *gin.Context) {
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

func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
func getTopPage(ctx *gin.Context) {
	ctx.HTML(200, "index.tmpl", gin.H{})
}
func getForm(ctx *gin.Context) {
	ctx.HTML(200, "form.tmpl", gin.H{})
}
func postForm(ctx *gin.Context) {
	//UserName := ctx.PostForm("User_name")
	ctx.JSON(http.StatusOK, gin.H{
		"name":     ctx.PostForm("User_name"),
		"mail":     ctx.PostForm("mail"),
		"age":    ctx.PostForm("age"),
		"gender":    ctx.PostForm("gender"),
		"onBusstop":  ctx.PostForm("board_bus_stop"),
		"offBusstop": ctx.PostForm("get_off_bus_stop"),
	})
}

func checkAnser 
