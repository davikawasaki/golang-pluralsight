package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	// fmt.Println(os.Environ())
}
