package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Printf("%v\n", arr)

	arrs := [3]string{"go", "is", "awesome"}
	fmt.Printf("%v\n", arrs)

	var arrB [2]bool
	arrB[0] = true
	arrB[1] = false

	mySlice := make([]int, 10)
	fmt.Printf("slice %v\n", mySlice)

	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	fmt.Println(myMap["one"])
}
