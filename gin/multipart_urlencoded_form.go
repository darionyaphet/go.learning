package main

import "github.com/gin-gonic/gin"

//curl -XPOST http://127.0.0.1:8989/form_post\?message\=m\&nick\=n
func main() {
	router := gin.Default()
	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")
		context.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8989")
}
