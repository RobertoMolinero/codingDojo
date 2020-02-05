package main

import (
	"fmt"
)

func main() {
	// Step 1: Single Person
	bob := "Bob"
	fmt.Println(Greet(bob))
}

//Greet Greets the persons that are transferred as parameters
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s.", name)
}