package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

//中间件使用
func main() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
	gin.DefaultErrorWriter = io.MultiWriter(file)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/test", func(context *gin.Context) {
		name := context.DefaultQuery("name", "default_name")
		panic("test panic")
		context.String(200, "%s", name)
	})
	r.Run()
}
