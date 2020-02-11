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
		limit float64
		want  float64
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

func TestSumOfGauss(t *testing.T) {
	tests := []struct {
		name  string
		limit float64
		want  float64
	}{
		{"Sum of Gauss 10", 10, 55},
		{"Sum of Gauss 100", 100, 5050},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfGauss(tt.limit); got != tt.want {
				t.Errorf("SumOfGauss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Least Common Multiple of 3 and 5", args{3, 5}, 15},
		{"Least Common Multiple of 4 and 6", args{4, 6}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeastCommonMultiple(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("LeastCommonMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
