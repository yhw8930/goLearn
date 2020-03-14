package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(context *gin.Context) {
		firstName := context.Query("first_name")
		lastName := context.DefaultQuery("last_name", "last_default_name")
		context.String(http.StatusOK, "%s,%s", firstName, lastName)
	})
	r.Run()
}
