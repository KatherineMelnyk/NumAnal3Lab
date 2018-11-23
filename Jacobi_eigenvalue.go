package main

import (
	"fmt"
	"math"
)

var iter int

func JacobiEigenvalue(A [][]float64) []float64 {
	var d [][]float64 //s, s1, s1t, temp
	var theta float64
	n := len(A)
	//s = EMatrix(n)

	for i := 0; i < len(A); i++ {
		d = append(d, []float64{})
		//	temp = append(temp, []float64{})
		for j := 0; j < n; j++ {
			d[i] = append(d[i], A[i][j])
			//		temp[i] = append(temp[i], 0.)
		}
	}
global:
	for {
		iter++
		i, j := maxOffDiagonal(d)

		theta = 0.5 * math.Atan2(2*d[j][i], d[i][i]-d[j][j])
		fmt.Printf("%v\n", theta)
		sin := math.Sin(theta)
		cos := math.Cos(theta)
		quadSin := math.Pow(sin, 2)
		quadCos := math.Pow(cos, 2)

		d[i][i] = quadCos*A[i][i] + 2*sin*cos*A[i][j] + quadSin*A[j][j]
		d[j][j] = -quadSin*A[i][i] - 2*sin*cos*A[i][j] + quadCos*A[j][j]
		d[i][j] = (quadCos-quadSin)*A[i][j] + cos*sin*(A[i][i]-A[j][j])
		d[j][i] = d[i][j]

		for k := 0; k < n; k++ {
			if k != i && k != j {
				d[i][k] = cos*A[i][k] + sin*A[j][k]
				//d[i][k] = -sin*A[j][k] + cos*A[i][k]
				d[k][i] = d[i][k]
			}
		}

		for k := 0; k < n; k++ {
			if k != i && k != j {
				d[j][k] = sin*A[i][k] - cos*A[j][k]
				//d[j][k] = cos*A[j][k] + sin*A[i][k]
				d[k][j] = d[j][k]
			}
		}

		for k := 0; k < n; k++ {
			for l := 0; l < n; l++ {
				if k != i && l != i && k != j && l != j {
					d[k][l] = A[k][l]
				}
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				A[i][j] = d[i][j]
			}
		}

		printMatrix(A)

		//s1 = givensRotation(theta, n, i, j)
		//s1t = T(s1)

		//temp = mul(s1t, d)

		//d = mul(temp, s1)

		//temp = mul(s, s1)

		//s = mul(s, s1)
		if iter == 3 {
			break
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					if math.Abs(d[i][j]) > EPS {
						continue global
					}
				}
			}
		}
		break
	}
	answer := eig(d)
	printVector(answer)
	return answer
}

func maxOffDiagonal(A [][]float64) (i, j int) {
	var max float64
	n := len(A)
	max = math.Abs(A[0][1])
	i = 0
	j = 1
	for p := 0; p < n; p++ {
		for q := 0; q < n; q++ {
			if p != q {
				if max < math.Abs(A[p][q]) {
					max = math.Abs(A[p][q])
					i = p
					j = q
				}
			}
		}
	}
	return i, j
}

func eig(A [][]float64) []float64 {
	var a []float64
	for i := 0; i < len(A); i++ {
		a = append(a, 0.)
	}
	for i := 0; i < len(A); i++ {
		a[i] = A[i][i]
	}
	return a
}

func givensRotation(theta float64, n, i, j int) [][]float64 {
	var res [][]float64
	res = EMatrix(n)
	res[i][i] = math.Cos(theta)
	res[j][j] = res[i][i]
	res[j][i] = math.Sin(theta)
	res[i][j] = -res[j][i]
	return res
}
