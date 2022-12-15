package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// フォームデータ用の構造体
type FormData struct {
	UserName    string
	MailAddress string
	Age         string
	Gender      string
	OnBusstop   string
	OffBusstop  string
	isNeedHelp  string
}

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
	var data FormData
	data.UserName = ctx.PostForm("User_name")
	data.MailAddress = ctx.PostForm("mail")
	data.Age = ctx.PostForm("age")
	data.Gender = ctx.PostForm("gender")
	data.OnBusstop = ctx.PostForm("board_bus_stop")
	data.OffBusstop = ctx.PostForm("get_off_bus_stop")
	data.isNeedHelp = ctx.PostForm("isNeedHelp")

	// 乗降バス停の重複を弾く
	if checkAnser(data.OnBusstop, data.OffBusstop) {
		ctx.HTML(200, "form.tmpl", gin.H{})

	} else {
		ctx.HTML(200, "result.tmpl",
			gin.H{
				"User_name":        data.UserName,
				"board_bus_stop":   data.OnBusstop,
				"get_off_bus_stop": data.OffBusstop,
				"isNeedHelp":       data.isNeedHelp,
			})
	}
}

func checkAnser(on string, off string) bool {
	return on == off
}
