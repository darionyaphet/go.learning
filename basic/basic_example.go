package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	const (
		a = iota
		b
		c
	)

	// a = 0 b = 1 c = 2
	fmt.Printf("a = %d b = %d c = %d \n", a, b, c)

	fmt.Println(add(1, 2))
}
