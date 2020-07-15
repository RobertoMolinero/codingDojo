package main

import (
	"fmt"
	"math"
)

func main() {
	number := 600851475143
	largestPrimeFactor := GetLargestPrimeFactor(number)
	fmt.Printf("LPF of %d: %d\n", number, largestPrimeFactor)
}

func GetLargestPrimeFactor(number int) int {
	candidate := number
	limit := int(math.Abs(math.Sqrt(float64(number))))

	for i := 2; i <= limit; i++ {
		if number%i == 0 {
			candidate = number / i
			return GetLargestPrimeFactor(candidate)
		}
	}

	return candidate
}
