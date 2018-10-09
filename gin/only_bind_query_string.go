package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func startPage(context *gin.Context) {
	var person Person
	if context.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(fmt.Sprintf("%s %s", person.Name, person.Address))
	}
	context.String(http.StatusOK, "Success")
}

func main() {
	route := gin.Default()
	route.Any("/bind", startPage)
	route.Run(":8989")
}
