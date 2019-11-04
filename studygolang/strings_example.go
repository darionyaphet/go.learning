package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	fmt.Println(strings.Count("five", "")) // before & after each rune

	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))

	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))

	fmt.Printf("%d\n", strings.IndexFunc("study golang", func(c rune) bool {
		if c > 'u' {
			return true
		}
		return false
	}))
}
