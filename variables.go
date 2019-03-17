package main

import (
	"fmt"
	"reflect"
)

var (
	name   = "Nigel"
	course = "Docker Deep Dive"
	module = 3.2
)

func main() {

	// Alternative way of declaring and initializing local variable in package levels
	// Go does not accept variable declaration that is not going to be used.
	// So course in here would emit the error: course declared and not used.
	//	name   := "Nigel"
	//	course := "Docker Deep Dive"
	//	module := 3.2
	ptr := &module

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	// fmt.Println("Memory address of *module* variable is ", &module)
	fmt.Println("Memory address of *module* variable is", &ptr,
		"and the value of *module* is", *ptr)
}
