package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run()
	//curl -H "Content-Type:application/json" -X POS"http://localhost:8080/testing" -d '{"name":"wang"}'
	//2006-01-02 15:04:05时间格式化
}

func testing(context *gin.Context) {
	var person Person
	//根据请求不同的Content—Type来做不同的binding操作
	if err := context.ShouldBind(&person); err == nil {
		context.String(200, "%v", person)
	} else {
		context.String(200, "person bind error:%v", err)
	}
}
