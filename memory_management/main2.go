package main

import (
	"debug/elf"
	"log"
)

func main() {
	f, err := elf.Open("/Users/darion.yaphet/go/src/github.com/darionyaphet/go.learning/memory_management/main")

	if err != nil {
		log.Fatal(err)
	}

	for _, section := range f.Sections {
		log.Println(section)
	}
}
