package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)

	timeUnix := time.Now().Unix()
	fmt.Println(timeUnix)

	timeUnixNano := time.Now().UnixNano()
	fmt.Println(timeUnixNano)

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	// int64 —>  string

	// string  —>  Time

	// string —>  int64
}
