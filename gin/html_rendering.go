package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl",
			gin.H{"title": "Main website"})
	})
	router.Run(":8989")
}
