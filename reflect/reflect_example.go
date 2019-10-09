package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

func main() {
	reflectType(3.14)     // type:float64 kind:float64
	reflectType(100)      // type:int kind:int
	reflectType(false)    // type:bool kind:bool
	reflectType("darion") // type:string kind:string
}
