package main

import "testing"

func TestEvenFibonacciNumbers(t *testing.T) {
	tests := []struct {
		name  string
		limit uint
		want  uint
	}{
		{"EvenFibonacciNumbers with limit 89", 89, 44},
		{"EvenFibonacciNumbers with limit 4M", 4000000, 4613732},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenFibonacciNumbers(tt.limit); got != tt.want {
				t.Errorf("EvenFibonacciNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    uint
		want uint
	}{
		{"Fibonacci Number 0", 0, 0},
		{"Fibonacci Number 1", 1, 1},
		{"Fibonacci Number 2", 2, 1},
		{"Fibonacci Number 3", 3, 2},
		{"Fibonacci Number 4", 4, 3},
		{"Fibonacci Number 5", 5, 5},
		{"Fibonacci Number 6", 6, 8},
		{"Fibonacci Number 7", 7, 13},
		{"Fibonacci Number 8", 8, 21},
		{"Fibonacci Number 9", 9, 34},
		{"Fibonacci Number 10", 10, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
