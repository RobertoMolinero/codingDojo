package main

import (
	"fmt"
	"strconv"
)

func main() {
	to100 := GenerateSequenceFrom1To100()
	fizz := Fizz(to100)

	for idx := range fizz {
		fmt.Printf("Number [%d]:", idx)
		fmt.Printf("\t")
		fmt.Print(fizz[idx])
		fmt.Printf("\n")
	}
}

//GenerateSequenceFrom1To100 Generates a list from 1 to 100
func GenerateSequenceFrom1To100() []int {
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

func Fizz(a []int) []string {
	s := make([]string, len(a))
	for idx, i := range a {
		if i%3 == 0 {
			s[idx] = fmt.Sprint("Fizz")
		} else {
			s[idx] = strconv.Itoa(i)
		}
	}
	return s
}
