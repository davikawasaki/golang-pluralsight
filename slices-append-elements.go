package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	// for index, data
	for _, i := range mySlice {
		fmt.Println("for range loop:", i)
	}

	// Append slice ITEMS (ellipsis) to other slice
	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)
}
