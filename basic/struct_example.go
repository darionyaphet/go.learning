package main

import "fmt"

type test struct{}

type Person struct {
	name string
	age  int
}

func Func(p *Person) {
	p.name = "darion.yaphet"
}

type human struct {
	sex int
}

type teacher struct {
	human
	name string
	age  int
}

type student struct {
	human
	name string
	age  int
}

func main() {
	a := Person{}
	a.name = "darion"
	a.age = 27
	fmt.Println(a)

	b := &Person{name: "darion", age: 27}
	fmt.Println(b)
	Func(b)
	fmt.Println(b)

	c := &struct {
		name string
		age  int
	}{name: "darion", age: 27}
	fmt.Println(c)

	t := teacher{name: "PHP", age: 15, human: human{sex: 0}}
	s := student{name: "GO", age: 8, human: human{sex: 1}}
	fmt.Println(t, s)

}
