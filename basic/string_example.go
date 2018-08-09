package main

import (
	"fmt"
	"strings"
)

func main() {

	// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("b", "a")) // 1
	fmt.Println("-----")

	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true
	fmt.Println("-----")

	fmt.Println(strings.ContainsAny("team", "i"))        // false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
	fmt.Println(strings.ContainsAny("foo", ""))          // false
	fmt.Println(strings.ContainsAny("", ""))             // false
	fmt.Println("-----")

	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
	fmt.Println("-----")

	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", ""))
	fmt.Println("-----")

	// are equal under Unicode case-folding.
	fmt.Println(strings.EqualFold("Go", "go")) // true
	fmt.Println("-----")

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	fmt.Println("-----")
}
