package main

import "fmt"

// use "const" (for constant)

// variable string
var language = "Go"

// empty variable (var [var name] [type data])
var empty string

// any variable inside function must be used)
func main() {
	// sort variable (only usable inside function)
	name := "Muhammad Yudha Prastya"

	// variable Numeric (int / integer)
	age := 21

	fmt.Println("Hi! i'm " + name)

	// use comma to combine numeric & string
	fmt.Println("i'm ", age)

	fmt.Println("i'm learning " + language)
}
