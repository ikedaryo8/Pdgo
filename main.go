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

/* ルーティング設定 */
// 接続テスト用
func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

// TOPページ
func getTopPage(ctx *gin.Context) {
	ctx.HTML(200, "index.tmpl", gin.H{})
}

// フォーム表示
func getForm(ctx *gin.Context) {
	ctx.HTML(200, "form.tmpl", gin.H{
		"message": "フォームを入力してください",
	})
}

// フォームが送信され
func postForm(ctx *gin.Context) {
	// フォームのデータをmap形式で扱う.
	// https://ashitani.jp/golangtips/tips_map.html#map_Map
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
	// 出力するものを判定
	//未入力フォームがある場合-> エラーメッセージ
	if flag { //未入力フォームを弾く
		ctx.HTML(200, "form.tmpl", gin.H{
			"message": "未入力フォームがあります",
		})
		//乗降バス停に同じバスを指定した場合-> エラーメッセージ
	} else if flag_selectSameBusstop {
		ctx.HTML(200, "form.tmpl", gin.H{
			"message": "同じ名前のバス停を指定することはできません.",
		})
	} else {
		ctx.HTML(200, "result.tmpl",
			gin.H{
				"User_name":        formData["UserName"],
				"board_bus_stop":   formData["OnBusstop"],
				"get_off_bus_stop": formData["OffBusstop"],
				"isNeedHelp":       formData["isNeedHelp"],
			})
		// ユーザーへメール送信
		// csvファイルに書き込む
	}
}

/* 入力判定関数 */
// バス停の重複を弾く
func isNotDuplicat(on string, off string) bool {
	return on == off
}

// 空の入力を弾く
func isNotNull(data string) bool {
	ans := true
	if data == " " {
		ans = false
	}
	return ans
}
