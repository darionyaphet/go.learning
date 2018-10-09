package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("welcome", func(context *gin.Context) {
		first := context.DefaultQuery("first", "Guest")
		last := context.Query("last")
		context.String(http.StatusOK, "Hello %s %s", first, last)
	})
}
