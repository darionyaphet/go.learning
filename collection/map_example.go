package main

import "fmt"

func main() {
	//var m map[string]int
	var m = make(map[string]int)
	fmt.Println(m)

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	m["one hundred"] = 100
	fmt.Println(m)

	var m0 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5, // Comma is necessary
	}

	fmt.Println(m0)
}
