package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.Must(template.New("hello").Parse("hello world"))
	t.Execute(os.Stdout, nil)

}
