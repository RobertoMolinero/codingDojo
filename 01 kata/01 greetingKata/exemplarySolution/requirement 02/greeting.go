package main

import (
	"fmt"
)

func main() {
	// Step 1: Single Person
	bob := "Bob"
	fmt.Println(Greet(&bob))
	// Step 2: Nil > Standard
	fmt.Println(Greet(nil))
}

//Greet Greets the persons that are transferred as parameters
func Greet(pointer *string) string {
	if pointer == nil {
		return fmt.Sprintf("Hello, my friend.")
	}
	return fmt.Sprintf("Hello, %s.", *pointer)
}
