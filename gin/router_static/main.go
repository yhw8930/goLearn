package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/fec", "./fec")
	r.Run()
}
