package main

import (
	"fmt"
)

func main() {
	// make(<type>, <len>, <cap>)
	myCourses := []string{"Docker", "Puppet", "Python"}
	// myCourses := make([]string, 5, 10)

	fmt.Printf("Length is: %d.\nCapacity is: %d",
		len(myCourses), cap(myCourses))

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice[4])

	mySlice[1] = 0
	fmt.Println(mySlice)

	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)
}
