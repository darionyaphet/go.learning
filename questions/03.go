package main

import "fmt"

// https://mp.weixin.qq.com/s/V00k8vfNanjYRAra5cviog

func funcMui(x, y int) (sum int, error) {
	return x + y, nil
}

func main() {
	s0 := make([]int, 5)
	s0 = append(s0, 1, 2, 3)
	fmt.Println(s0)

	s1 := make([]int, 0)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)

	fmt.Println(funcMui(1, 2))
}
