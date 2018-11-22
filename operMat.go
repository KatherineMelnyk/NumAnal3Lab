package main

import "math"

func mulMatVec(A [][]float64, b []float64) []float64 {
	var tmp, k []float64
	n := len(A)
	for i := 0; i < n; i++ {
		tmp = append(tmp, 0)
		k = append(k, 0)
	}
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			//noinspection ALL
			tmp[i] += A[i][k] * b[k]
		}
	}

	return tmp
}

func eIt(x []float64) []float64 {
	var e []float64
	for i := 0; i < len(x); i++ {
		e = append(e, 0.)
	}
	for j := 0; j < len(x); j++ {
		e[j] = x[j] / vecNorm(x)
	}
	return e
}

func vecNorm(v []float64) float64 {
	var res float64
	for i := 0; i < len(v); i++ {
		res += math.Pow(v[i], 2)
	}
	return math.Sqrt(res)
}

func scalar(x, m []float64) float64 {
	var prod float64
	for i := 0; i < len(x); i++ {
		prod += x[i] * m[i]
	}
	prod = math.Sqrt(prod)
	return prod
}

func normMatrix(A [][]float64) float64 {
	var sum float64
	max := 0.
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if math.Abs(A[j][i]) > max {
				max = math.Abs(A[j][i])
			}
		}
		sum += max
		max = 0
	}
	return sum
}

func EMatrix(n int) [][]float64 {
	var E [][]float64
	for i := 0; i < n; i++ {
		E = append(E, []float64{})
		for j := 0; j < n; j++ {
			if i == j {
				E[i] = append(E[i], 1)
			} else {
				E[i] = append(E[i], 0)
			}
		}
	}

	return E
}

func mulMatSc(A [][]float64, s float64) [][]float64 {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			A[i][j] *= s
		}
	}
	return A
}

func minusMat(A, B [][]float64) [][]float64 {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			A[i][j] -= B[i][j]
		}
	}
	return A
}

func T(A [][]float64) [][]float64 {
	n := len(A)
	var B [][]float64
	for i := 0; i < n; i++ {
		B = append(B, []float64{})
		for j := 0; j < n; j++ {
			B[i] = append(B[i], 0.)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			B[j][i] = A[i][j]
		}
	}
	return B
}

func mul(A, B [][]float64) [][]float64 {
	var tmp, k [][]float64
	n := len(A)
	for i := 0; i < n; i++ {
		tmp = append(tmp, []float64{})
		k = append(k, []float64{})
		for j := 0; j < n; j++ {
			tmp[i] = append(tmp[i], 0)
			k[i] = append(k[i], 0)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				tmp[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return tmp
}
