package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "bye" {
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
