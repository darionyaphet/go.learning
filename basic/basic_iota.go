package main

import "fmt"

// https://splice.com/blog/iota-elegant-constants-golang/
// https://yourbasic.org/golang/iota/

func main() {
	//Iota basic example
	const (
		C0 = iota
		C1 = iota
		C2 = iota
	)

	const (
		C00 = iota
		C01
		C02
	)

	//Start from one
	const (
		C001 = iota + 1
		C002
		C003
	)

	//Skip value
	const (
		C0001 = iota + 1
		_
		C0003
		C0004
	)

	// Complete enum type with strings [best practice]

	type Direction int //create a new integer type,

	const (
		North Direction = iota
		East
		South
		West
	)

	var d Direction = North
	fmt.Print(d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}

	fmt.Println(C0, C1, C2)          // "0 1 2"
	fmt.Println(C00, C01, C02)       // "0 1 2"
	fmt.Println(C001, C002, C003)    // "1 2 3"
	fmt.Println(C0001, C0003, C0004) // "1 3 4"
}
