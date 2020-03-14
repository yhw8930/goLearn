package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./gin/other_template/template/*")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"title": "test.html",
		})
	})
	r.Run()
}
