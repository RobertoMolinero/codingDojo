package main

//Point Defines a point in the grid
type Point struct {
	X int
	Y int
}

//LatticePathsDynamic Moves the nodes from the end point back to the origin and calculates the path possibilities from the predecessors.
func LatticePathsDynamic(gridSize int) int {
	m := InitializeEdges(gridSize)

	for i := gridSize - 1; i >= 0; i-- {
		for j := gridSize - 1; j >= 0; j-- {
			m[Point{i, j}] = m[Point{i + 1, j}] + m[Point{i, j + 1}]
		}
	}

	return m[Point{0, 0}]
}

//InitializeEdges Generates a directory with the values for the limit values of the grid.
func InitializeEdges(gridSize int) map[Point]int {
	m := make(map[Point]int)

	for i := 0; i <= gridSize; i++ {
		m[Point{gridSize, i}] = 1
		m[Point{i, gridSize}] = 1
	}

	m[Point{gridSize, gridSize}] = 0

	return m
}
