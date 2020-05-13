package main

import "fmt"

type person struct {
	name string
	age  int
}

type name string

func main() {
	p := person{
		name: "darion",
		age:  18,
	}

	fmt.Println(p.name)
}
