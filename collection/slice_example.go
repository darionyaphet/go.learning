package main

import "fmt"

func main() {

	// 声明一个未指定大小的数组来定义切片  var identifier []type
	// 或使用make()函数来创建切片        var s []type = make([]type, len)

	primes := [6]int{0, 1, 2, 3, 4, 5}
	var s []int = primes[1:4]
	fmt.Println(s)

	s0 := []int{1, 2, 3}
	fmt.Println(s0)

	s1 := make([]int, 10, 10)
	fmt.Println(s1)
	fmt.Printf("slice=%v len=%d cap=%d \n", s1, len(s1), cap(s1))

	var nilS []int
	fmt.Println(nilS)
	if nilS == nil {
		fmt.Printf("nilS == nill")
	}

	// 追加空切片
	var numbers []int
	numbers = append(numbers, 0)
	fmt.Println(numbers)

	// 二维数组
	dimension := [][]string{{"00", "01"}, {"10", "11"}, {"20", "21"}, {"30", "31"}, {"40", "41"}}
	for i := 0; i < len(dimension); i++ {
		for j := 0; j < len(dimension[0]); j++ {
			fmt.Printf("dimension[%d][%d] = %s\n", i, j, dimension[i][j])
		}
	}
}
