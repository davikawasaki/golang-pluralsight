package main

import (
	"fmt"
)

func main() {
	name := "Nigel"
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course)
	changeCourseByVal(course)
	fmt.Println("\nYou are now watching old course", course)
	changeCourseByRef(&course)
	fmt.Println("\nYou are now watching updated course", course)
}

func changeCourseByRef(course *string) string {
	*course = "First Look: Native Docker Clustering"
	fmt.Println("Trying to change your course to", *course)
	return *course
}

func changeCourseByVal(course string) string {
	course = "First Look: Native Docker Clustering"
	fmt.Println("Trying to change your course to", course)
	return course
}
