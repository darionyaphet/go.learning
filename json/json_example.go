package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Book struct {
		Name  string
		Price float64 // `json:"price,string"`
	}

	var person = struct {
		Name string
		Age  int
		Book
	}{
		Name: "polaris",
		Age:  30,
		Book: Book{
			Price: 3.4,
			Name:  "Go语言",
		},
	}

	buf, _ := json.Marshal(person)
	fmt.Println(string(buf))
}
