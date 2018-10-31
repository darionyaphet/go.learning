package main

//https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/
type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}
