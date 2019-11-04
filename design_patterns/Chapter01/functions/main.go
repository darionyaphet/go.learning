package main

import (
	"errors"
	"fmt"
)

func hello(message string) error {
	fmt.Printf("Hello %s\n", message)
	return nil
}

func doesReturnError() error {
	err := errors.New("this function simply returns an error")
	return err
}

func sum(args ...int) (result int) {
	for _, v := range args {
		result += v
	}
	return
}

func main() {
	err := errors.New("Error example")
	fmt.Printf("%v\n", err)

	fmt.Printf("%d\n", sum(1, 2, 3))
	fmt.Printf("%d\n", sum(4, 5, 6, 7, 8))

	e := doesReturnError()
	if e != nil {
		//panic(e)
	}
}
