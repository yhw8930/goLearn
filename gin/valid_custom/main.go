package main

import (
	"time"

	"gopkg.in/go-playground/validator.v9"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

//自定义验证规则
func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	r.GET("/bookable", func(context *gin.Context) {
		var b Booking
		if err := context.ShouldBind(&b); err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(200, gin.H{"message": "ok!", "booking": b})
	})
	r.Run()
	//curl -X GET "http://localhost:8080/bookable?check_in=2020-03-14&check_out=2020-03-26"
}
