package main

//https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type alphaReader struct {
	source  string
	current int
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{source: src}
}

func alpha(b byte) byte {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return b
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	if a.current >= len(a.source) {
		return 0, io.EOF
	}

	x := len(a.source) - a.current
	n, bound := 0, 0
	if x >= len(p) {
		bound = len(p)
	} else if x <= len(p) {
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		if char := alpha(a.source[a.current]); char != 0 {
			buf[n] = char
		}
		n++
		a.current++
	}
	copy(p, buf)
	return n, nil
}

func main() {
	reader := strings.NewReader("Altissima quaeque flumina minimo sono labi")
	p := make([]byte, 9)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n])) //should handle any remainding bytes.
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
	}
}
