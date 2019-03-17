package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	// Ways of declaring new instance of struct with 0 value
	//	var DockerDeepDive courseMeta
	//	DockerDeepDive := new(courseMeta)  // Gives a pointer with new function

	// Ways of declaring new instance of struct with starting values
	DockerDeepDive := courseMeta{
		Author: "Davi Kawasaki",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println("\nDocker Deep Dive author is:", DockerDeepDive.Author)
	fmt.Println("\nDocker Deep Dive rating is:", DockerDeepDive.Rating)
}
