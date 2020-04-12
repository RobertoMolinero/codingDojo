package main

//LatticePathsCombinatorial Calculates the number of paths in the grid using the binomial coefficient.
func LatticePathsCombinatorial(gridSize int) int {
	paths := 1

	for i:= 0; i < gridSize; i++ {
		paths *= (2 * gridSize) - i
		paths /= i + 1
	}

	return paths
}
