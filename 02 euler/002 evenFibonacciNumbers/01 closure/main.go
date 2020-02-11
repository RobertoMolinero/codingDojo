package main

import "fmt"

func main() {
	sum := EvenFibonacciNumbers(4000000)
	fmt.Printf("Sum: %d\n", sum)
}

func EvenFibonacciNumbers(limit int) int {
	var sum int
	f := Fibonacci()
	for {
		fib := f()
		if fib > limit {
			break
		}
		sum += fib
		f()
		f()
	}
	return sum
}

func Fibonacci() func() int {
	current, next := 0, 1
	return func() int {
		defer func() {
			tmp := current
			current, next = next, tmp+next
		}()
		return current
	}
}
