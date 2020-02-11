package main

import "fmt"

func main() {
	sum := EvenFibonacciNumbers(4000000)
	fmt.Printf("Sum: %d\n", sum)
}

func EvenFibonacciNumbers(limit uint) uint {
	var sum uint
	var i uint

	for {
		fib := Fibonacci(i)
		if fib > limit {
			break
		}
		sum += fib
		i += 3
	}

	return sum
}

func Fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}

	var n2, n1 uint = 0, 1

	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}
