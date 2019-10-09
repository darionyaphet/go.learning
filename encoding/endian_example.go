package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var num uint64
	num = 0x1234
	fmt.Printf("num = %x\n", num)
	enc := make([]byte, 8)

	binary.BigEndian.PutUint64(enc, num)
	fmt.Printf("Big Endian    = %x\n", enc)

	binary.LittleEndian.PutUint64(enc, num)
	fmt.Printf("Little Endian = %x\n", enc)
}
