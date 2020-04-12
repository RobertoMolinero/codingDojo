package main

import (
	"reflect"
	"testing"
)

func TestCreateGrid(t *testing.T) {

	mapOfTwo := make(map[Point]int)
	mapOfTwo[Point{2, 0}] = 1
	mapOfTwo[Point{2, 1}] = 1
	mapOfTwo[Point{0, 2}] = 1
	mapOfTwo[Point{1, 2}] = 1
	mapOfTwo[Point{2, 2}] = 0

	mapOfThree := make(map[Point]int)
	mapOfThree[Point{3, 0}] = 1
	mapOfThree[Point{3, 1}] = 1
	mapOfThree[Point{3, 2}] = 1
	mapOfThree[Point{0, 3}] = 1
	mapOfThree[Point{1, 3}] = 1
	mapOfThree[Point{2, 3}] = 1
	mapOfThree[Point{3, 3}] = 0

	tests := []struct {
		name     string
		gridSize int
		want     map[Point]int
	}{
		{"Map of Two", 2, mapOfTwo},
		{"Map of Two", 3, mapOfThree},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeEdges(tt.gridSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeEdges() = %v, want %v", got, tt.want)
				t.Errorf("InitializeEdges() = %v, want %v", len(got), 2*(tt.gridSize-1)+1)
			}
		})
	}
}

func TestLatticePathsDynamic(t *testing.T) {

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
			if got := LatticePathsDynamic(tt.gridSize); got != tt.want {
				t.Errorf("LatticePathsDynamic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLatticePathsDynamicWithTable(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	benchmarks := []struct {
		name string
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
				LatticePathsDynamic(bm.gridSize)
			}
		})
	}
}
