package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Sum: %v\n", MultiplesOf3And5(1000))
}

func MultiplesOf3And5(limit float64) float64 {
	sum3 := multiples(3, limit)
	sum5 := multiples(5, limit)
	sumLcm := multiples(float64(leastCommonMultiple(3, 5)), limit)
	return sum3 + sum5 - sumLcm
}

func multiples(of float64, limit float64) float64 {
	factor := math.Trunc((limit - 1) / of)
	sum := sumOfGauss(factor)
	return of * sum
}

func sumOfGauss(limit float64) float64 {
	return 0.5 * limit * (limit + 1)
}

func leastCommonMultiple(m int, n int) int {
	lcm := 0
	if m > n {
		lcm = m
	} else {
		lcm = n
	}

	for {
		if lcm%m == 0 && lcm%n == 0 {
			return lcm
		}
		lcm++
	}
}
