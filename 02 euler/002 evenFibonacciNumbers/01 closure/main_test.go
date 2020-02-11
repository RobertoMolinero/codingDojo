package main

import (
	"testing"
)

func TestEvenFibonacciNumbers(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  int
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
