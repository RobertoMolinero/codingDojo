package main

import (
	"fmt"
	"strings"
)

func main() {
	// Step 1: Single Person
	bob := "Bob"
	fmt.Println(Greet(&bob))
	// Step 2: Nil > Standard
	fmt.Println(Greet(nil))
	// Step 3: Shouting
	jerry := "JERRY"
	fmt.Println(Greet(&jerry))
}

//Greet Greets the persons that are transferred as parameters
func Greet(pointer *string) string {
	if pointer == nil {
		return fmt.Sprintf("Hello, my friend.")
	}
	name := *pointer
	if name == strings.ToUpper(name) {
		return fmt.Sprintf("HELLO %s!", name)
	}
	return fmt.Sprintf("Hello, %s.", name)
}
