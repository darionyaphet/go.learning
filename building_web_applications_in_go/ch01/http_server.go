package main

import (
	"log"
	"net/http"
)

// https://github.com/unknwon/building-web-applications-in-go/blob/master/articles/01.md
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world!"))
	})

	log.Println("Starting HTTP server...") // 启动 HTTP 服务器
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
