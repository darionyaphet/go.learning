package main

// https://studygolang.com/articles/443

import "fmt"

type Painter interface {
	Draw()
}

type Cowboy interface {
	Draw()
}

type Xiaowang struct {
}

type XiaowangAsPainter Xiaowang

func (p *XiaowangAsPainter) Draw() {
	fmt.Println("I'm painting.")
}

type XiaowangAsCowboy Xiaowang

func (p *XiaowangAsCowboy) Draw() {
	fmt.Println("I'm drawing.")
}

func (Xiaowang) Draw() {
	fmt.Println("i am drawing a paper.")
}

func main() {
	var xw Xiaowang
	var painter Painter = xw
	painter.Draw()

	var painter2 Painter = &xw
	painter2.Draw()

	painter3 := Painter(xw)
	painter3.Draw()

	painter4 := Painter(&xw)
	painter4.Draw()

	var p Painter = (*XiaowangAsPainter)(&xw)
	p.Draw()

	var c Cowboy = (*XiaowangAsCowboy)(&xw)
	c.Draw()

}
