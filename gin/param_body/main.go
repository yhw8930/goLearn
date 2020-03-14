package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(context *gin.Context) {
		all, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			context.Abort()
		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(all))
		firstName := context.PostForm("first_name")
		lastName := context.DefaultPostForm("last_name", "last_default_name")
		context.String(http.StatusOK, "%s,%s,%s", firstName, lastName, string(all))

	})
	r.Run()
	//curl -X POST "http://localhost:8080/test" -d 'first_name=wang&last_name=kai'
}
