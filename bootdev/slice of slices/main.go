package main

func createMatrix(rows, cols int) [][]int {

	matrix := make([][]int, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			cellComputation := i * j
			matrix[i][j] = cellComputation
		}
	}

	return matrix
}
