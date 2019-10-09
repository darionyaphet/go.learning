package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	buf := new(bytes.Buffer)
	pi := math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x\n", buf.Bytes())

	var piVar float64
	piBuffer := bytes.NewReader(buf.Bytes())

	binary.Read(piBuffer, binary.LittleEndian, &piVar)
	fmt.Println("pi", piVar) // pi   3.141592653589793
}
