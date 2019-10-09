package main

import "fmt"

type dialOption struct {
	Username string
	Password string
	Service  string
}

type DialOption interface {
	apply(*dialOption)
}

type funcOption struct {
	f func(*dialOption)
}

func (fdo *funcOption) apply(do *dialOption) {
	fdo.f(do)
}

func newFuncOption(f func(*dialOption)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func withUserName(s string) DialOption {
	return newFuncOption(func(o *dialOption) {
		o.Username = s
	})
}

func withPasswordd(s string) DialOption {
	return newFuncOption(func(o *dialOption) {
		o.Password = s
	})
}

func withService(s string) DialOption {
	return newFuncOption(func(o *dialOption) {
		o.Service = s
	})
}

func defaultOptions() dialOption {
	return dialOption{
		Service: "test",
	}
}

type clientConn struct {
	timeout int
	dopts   dialOption
}

func NewClient(address string, opts ...DialOption) {
	cc := &clientConn{
		timeout: 30,
		dopts:   defaultOptions(),
	}

	for _, opt := range opts {
		opt.apply(&cc.dopts)
	}

	fmt.Printf("%+v\n", cc.dopts)
}

func main() {
	NewClient("127.0.0.1", withPasswordd("654321"), withService("habox"))
	NewClient("127.0.0.1", withService("habox"))
}
