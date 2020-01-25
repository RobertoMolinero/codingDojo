package main

func main() {

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
