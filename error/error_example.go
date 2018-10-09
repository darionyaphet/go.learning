package main

import "fmt"

type Err struct {
	err string
}

func (e *Err) Error() string {
	return e.err
}

func makeErr() *Err {
	return nil
}

func main() {
	var err error = makeErr()
	fmt.Println(err, err != nil)
}
