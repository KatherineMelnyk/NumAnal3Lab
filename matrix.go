package main

import "fmt"

func MyVec(n int) []float64 {
	var vector []float64
	for i := 0; i < n; i++ {
		vector = append(vector, 1.)
	}
	return vector
}

func MyMatrix(n int) [][]float64 {
	var matrix [][]float64
	for i := 0; i < n; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j < n; j++ {
			if i == j && i != 0 && j != 0 {
				matrix[i] = append(matrix[i], 0.)
			} else if i == 0 && j == 0 && i == j {
				matrix[i] = append(matrix[i], 1.)
			} else {
				matrix[i] = append(matrix[i], float64(j+1))
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				break
			} else {
				matrix[i][j] *= -1
			}
		}
	}
	return matrix
}

func symmetricM(n int) [][]float64 {
	var matrix [][]float64
	for i := 0; i < n; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j < n; j++ {
			matrix[i] = append(matrix[i], 0.)
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				matrix[i][j] = 1.
			} else {
				matrix[i][j] = float64(j)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[j][i] = matrix[i][j]
		}
	}
	return matrix
}

func printVector(b []float64) {
	for i := range b {
		fmt.Printf("  %14.8f", b[i])
		fmt.Print("\n")
	}
	fmt.Printf("\n")
}

func printMatrix(A [][]float64) {
	for i := range A {
		for j := range A[i] {
			fmt.Printf("%13.8f", A[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Printf("\n")
}
