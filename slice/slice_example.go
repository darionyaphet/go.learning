package main

import "fmt"

var i int

func fa(s []int, n int) int {
	i = n
	for i = 0; i < len(s); i++ {

	}
	return i
}

func fb(s []int, n int) int {
	i = n
	for i = range s {
	}
	return i
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(fa(s, -1), fb(s, -1))

	s = nil
	fmt.Println(fa(s, -1), fb(s, -1))
}
