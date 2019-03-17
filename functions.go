package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "davi kawasaki"

	//	fmt.Println(converter(module, author))

	// Anonymous functions are self-executed
	func() {
		fmt.Println(converter(module, author))
	}
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}
