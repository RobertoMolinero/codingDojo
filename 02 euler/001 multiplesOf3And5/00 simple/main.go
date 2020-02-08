package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Sum: %d\n", MultiplesOf3And5(1000))
}

func MultiplesOf3And5(limit int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
