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
	// Step 4: Pair
	jill := "Jill"
	jane := "Jane"
	fmt.Println(Greet(&jill, &jane))
}

//Greet Greets the persons that are transferred as parameters
func Greet(pointer ...*string) string {
	if isNil, standardGreeting := checkForNil(pointer); isNil {
		return standardGreeting
	}
	if len(pointer) == 1 {
		name := *pointer[0]
		return greetSinglePerson(name)
	}

	var listOfNames []string
	for i := 0; i < len(pointer); i++ {
		s := *pointer[i]
		listOfNames = append(listOfNames, s)
	}
	return greetListOfPersons(listOfNames)
}

func checkForNil(pointer []*string) (bool, string) {
	if pointer == nil || pointer[0] == nil {
		return true, fmt.Sprintf("Hello, my friend.")
	}
	return false, ""
}

func greetSinglePerson(name string) string {
	if name == strings.ToUpper(name) {
		return fmt.Sprintf("HELLO %s!", name)
	}
	return fmt.Sprintf("Hello, %s.", name)
}

func greetListOfPersons(listOfNames []string) string {
	var s string

	s += fmt.Sprint("Hello, ")
	for idx, name := range listOfNames {
		s += fmt.Sprintf("%s", name)
		if idx < (len(listOfNames) - 1) {
			s += fmt.Sprint(" and ")
		} else {
			s += fmt.Sprint(".")
		}
	}

	return s
}
