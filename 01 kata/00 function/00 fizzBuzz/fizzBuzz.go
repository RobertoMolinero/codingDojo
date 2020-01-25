package main

import "fmt"

func main() {
	to100 := generateSequenceFrom1To100()
	f := fizz(to100)

	for idx, i := range f {
		fmt.Printf("Number [%d]: %s\n", idx, i)
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

	t := make([]string, len(s))
	for idx, i := range s {
		if i%3 == 0 {
			t[idx] = "fizz"
		} else {
			t[idx] = fmt.Sprint(i)
		}
	}

	return t
}
