package main

import "fmt"

type Handler interface {
	Work(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Work(k, v)
		}
	}
}

type welcome string

func (w welcome) Work(k, v interface{}) {
	fmt.Printf("%s, name: %s age: %d\n", w, k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["a"] = 20
	persons["b"] = 23
	persons["c"] = 26

	var w welcome = "welcome"
	Each(persons, w)
}
