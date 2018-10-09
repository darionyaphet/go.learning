package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func GetOrElse(err error, value interface{}, defaultValue interface{}) interface{} {
	if err == nil {
		return value
	} else {
		return defaultValue
	}
}

func String2Int(value string, defaultValue int) int {
	val, err := strconv.Atoi(value)
	return GetOrElse(err, val, defaultValue).(int)
}

func String2Int64(value string, defaultValue int64) int64 {
	val, err := strconv.ParseInt(value, 10, 64)
	return GetOrElse(err, val, defaultValue).(int64)
}

func Int2String(value int) string {
	return strconv.Itoa(value)
}

func Int642String(value int64) string {
	return strconv.FormatInt(value, 10)
}

func String2Float32(value string, defaultValue float32) float32 {
	val, err := strconv.ParseFloat(value, 32)
	return GetOrElse(err, val, defaultValue).(float32)
}

func String2Float64(value string, defaultValue float64) float64 {
	val, err := strconv.ParseFloat(value, 64)
	return GetOrElse(err, val, defaultValue).(float64)
}

//func Float322String(value float32, defaultValue string) float32 {
//
//}
//
//func Float642String(value float64, defaultValue string) float64 {
//
//}

func Slice2String() {

}

func Map2String(m map[string]string) string {
	return ""
}

func Map2JSON(m map[string]int) string {
	array, _ := json.Marshal(m)
	return string(array)
}

func main() {
	fmt.Println(String2Int("8989", -1))
	fmt.Println(String2Int("abcd", -1))

	fmt.Println(String2Int64("8989", -1))
	fmt.Println(String2Int64("abcd", -1))

	fmt.Println(Int2String(8989))
	fmt.Println(Int642String(8989))
	//
	//fmt.Println(String2Float32("3.14", 0.618))
	//fmt.Println(String2Float32("a.bc", 0.618))
	//
	//fmt.Println(String2Float64("3.14", 0.618))
	//fmt.Println(String2Float64("a.bc", 0.618))

	m0 := make(map[string]int)
	m0["1"] = 1
	m0["2"] = 2
	m0["3"] = 3
	m0["4"] = 4
	m0["5"] = 5
	fmt.Println(Map2JSON(m0))
	fmt.Println()
}
