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
	// Step 5: List
	amy := "Amy"
	brian := "Brian"
	charlotte := "Charlotte"
	fmt.Println(Greet(&amy, &brian, &charlotte))
	// Step 6: Mixing
	brianUppercase := "BRIAN"
	fmt.Println(Greet(&amy, &brianUppercase, &charlotte))
	// Step 7: String with comma
	charlieDianne := "Charlie, Dianne"
	fmt.Println(Greet(&bob, &charlieDianne))
	// Step 8: Escape
	charlieDianneEscape := "\"Charlie, Dianne\""
	fmt.Println(Greet(&bob, &charlieDianneEscape))
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

		if strings.Contains(s, ",") {
			if strings.Contains(s, "\"") {
				listOfNames = append(listOfNames, strings.Replace(s, "\"", "", -1))
			} else {
				split := strings.Split(s, ",")
				for _, n := range split {
					listOfNames = append(listOfNames, strings.TrimSpace(n))
				}
			}
		} else {
			listOfNames = append(listOfNames, s)
		}
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
	var greet []string
	var shout []string
	for _, name := range listOfNames {
		if name == strings.ToUpper(name) {
			shout = append(shout, name)
		} else {
			greet = append(greet, name)
		}
	}
	s += fmt.Sprint("Hello, ")
	for idx, name := range greet {
		s += fmt.Sprintf("%s", name)

		if idx < (len(greet) - 2) {
			s += fmt.Sprint(", ")
		} else if idx < (len(greet) - 1) {
			s += fmt.Sprint(" and ")
		} else {
			s += fmt.Sprint(". ")
		}
	}

	for _, name := range shout {
		s += fmt.Sprintf("AND HELLO %s!", name)
	}

	return strings.TrimSpace(s)
}
