package main

import (
	"fmt"
)

func main() {
	to100 := generateSequenceFrom1To100()
	fb := fizzBuzz(to100)

	for idx, i := range fb {
		fmt.Printf("Number [%d]:", idx+1)
		fmt.Printf("\t")
		fmt.Print(i)
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

func fizzBuzz(a []int) []string {
	s := make([]string, len(a))
	for idx, i := range a {
		if i%3 == 0 {
			s[idx] += fmt.Sprint("Fizz")
		}
		if i%5 == 0 {
			s[idx] += fmt.Sprint("Buzz")
		}
		if len(s[idx]) == 0 {
			s[idx] = fmt.Sprint(i)
		}
	}
	return s
}
