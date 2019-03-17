package main

import (
	"fmt"
)

func main() {
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	// random loop start offset
	for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is: %v", key, value)
	}

	// fmt.Println(testMap["C"])

	testMap["A"] = 100
	testMap["F"] = 1973
	fmt.Println(testMap)

	delete(testMap, "F")
	fmt.Println(testMap)

	// Maps are passed to functions by reference (reference types)
	// Unsafe for concurrency
	// Big maps: specify size for large maps e.g. make(map[<keyType>]<valueType>, size)
	// Implementation of hash tables
}
