package main

import "testing"

func TestLatticePathsCombinatorial(t *testing.T) {

	tests := []struct {
		name     string
		gridSize int
		want     int
	}{
		{"Grid of Two", 2, 6},
		{"Grid of Two", 3, 20},
		{"Grid of Two", 4, 70},
		{"Grid of Two", 5, 252},
		{"Grid of Two", 6, 924},
		{"Grid of Two", 7, 3432},
		{"Grid of Two", 20, 137846528820},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LatticePathsCombinatorial(tt.gridSize); got != tt.want {
				t.Errorf("LatticePathsCombinatorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLatticePathsCombinatorial(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	benchmarks := []struct {
		name     string
		gridSize int
	}{
		{"gridSize=001", 1},
		{"gridSize=002", 2},
		{"gridSize=005", 5},
		{"gridSize=010", 10},
		{"gridSize=020", 20},
		{"gridSize=050", 50},
		{"gridSize=100", 100},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LatticePathsCombinatorial(bm.gridSize)
			}
		})
	}
}
