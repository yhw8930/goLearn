package main

import "github.com/gin-gonic/gin"

type Person struct {
	Age     int    `form:"age" binding:"required,gt=10"`
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBind(&person); err != nil {
			context.String(500, "%v", err)
			context.Abort()
			return
		}
		context.String(200, "%v", person)
	})
	r.Run()
}
