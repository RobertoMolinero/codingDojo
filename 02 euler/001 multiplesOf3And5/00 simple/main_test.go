package main

import "testing"

func BenchmarkMultiplesOf3And5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MultiplesOf3And5(1000)
	}
}

func TestMultiplesOf3And5(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  int
	}{
		{"MultiplesOf3And5 with Limit 10", 10, 23},
		{"MultiplesOf3And5 with Limit 100", 100, 2318},
		{"MultiplesOf3And5 with Limit 1000", 1000, 233168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplesOf3And5(tt.limit); got != tt.want {
				t.Errorf("MultiplesOf3And5() = %v, want %v", got, tt.want)
			}
		})
	}
}
