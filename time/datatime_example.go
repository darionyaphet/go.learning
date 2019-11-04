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

	//https://www.toutiao.com/a6740498459050115596/
	now := time.Now()
	fmt.Printf("current time: %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()

	fmt.Printf("current timestamp : %v\n", timeStamp1)
	fmt.Printf("current timestamp : %v\n", timeStamp2)
}
