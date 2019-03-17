package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 1, 4)
	fmt.Printf("Length is: %d Capacity is: %d", len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ {
		// append normally until reaches full capacity.
		// when that happens, it doubles the array capacity.
		// e.g. 4 becomes 8, 8 becomes 16, and so on...
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", cap(mySlice))
	}
}
