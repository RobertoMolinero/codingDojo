package main

import "fmt"

func main() {
	to100 := generateSequenceFrom1To100()

	resultOfFizz := fizz(to100)
	resultOfBuzz := buzz(to100)

	for idx, i := range to100 {
		fmt.Printf("Number [%d]:", i)
		fmt.Printf("\t")

		if len(resultOfFizz[idx])+len(resultOfBuzz[idx]) > 0 {
			fmt.Printf("%s", resultOfFizz[idx])
			fmt.Printf("%s", resultOfBuzz[idx])
		} else {
			fmt.Print(i)
		}

		fmt.Printf("\n")
	}
}

func generateSequenceFrom1To100() []int {
	start := 1
	max := 100
	step := 1
	s := make([]int, max)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

func fizz(s []int) []string {
	return substituteMultiples(s, 3, "Fizz")
}

func buzz(s []int) []string {
	return substituteMultiples(s, 5, "Buzz")
}

func substituteMultiples(s []int, factor int, word string) []string {
	t := make([]string, len(s))
	for idx, i := range s {
		if i%factor == 0 {
			t[idx] = word
		}
	}
	return t
}
