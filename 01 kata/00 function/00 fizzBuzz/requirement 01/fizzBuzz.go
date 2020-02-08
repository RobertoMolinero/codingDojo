package main

import "fmt"

func main() {
	to100 := GenerateSequenceFrom1To100()
	for idx, i := range to100 {
		fmt.Printf("Number [%d]:", idx)
		fmt.Printf("\t")
		fmt.Print(i)
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
