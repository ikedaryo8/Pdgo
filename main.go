package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	ctx.HTML(200, "form.tmpl", gin.H{
		"message": "フォームを入力してください",
	})
}

// https://ashitani.jp/golangtips/tips_map.html#map_Map
func postForm(ctx *gin.Context) {
	formData := map[string]string{
		"UserName":    ctx.PostForm("User_name"),
		"MailAddress": ctx.PostForm("mail"),
		"Age":         ctx.PostForm("age"),
		"Gender":      ctx.PostForm("gender"),
		"OnBusstop":   ctx.PostForm("board_bus_stop"),
		"OffBusstop":  ctx.PostForm("get_off_bus_stop"),
		"isNeedHelp":  ctx.PostForm("isNeedHelp"),
	}

	// 乗降バス停の重複を弾く
	flag_selectSameBusstop := isNotDuplicat(formData["OnBusstop"], formData["OffBusstop"])
	flag_DataNotNull := true
	for key := range formData {
		flag_DataNotNull = flag_DataNotNull && isNotNull(formData[key])
	}

	flag := flag_selectSameBusstop && flag_DataNotNull
	if flag_selectSameBusstop {
		ctx.HTML(200, "form.tmpl", gin.H{
			"message": "同じ名前のバス停を指定することはできません.",
		})
	} else if flag {
		ctx.HTML(200, "form.tmpl", gin.H{
			"message": "値が不正です",
		})

	} else {
		ctx.HTML(200, "result.tmpl",
			gin.H{
				"User_name":        formData["UserName"],
				"board_bus_stop":   formData["OnBusstop"],
				"get_off_bus_stop": formData["OffBusstop"],
				"isNeedHelp":       formData["isNeedHelp"],
			})
	}
}

func isNotDuplicat(on string, off string) bool {
	return on == off
}

func isNotNull(data string) bool {
	ans := true
	if data == " " {
		ans = false
	}
	return ans
}
