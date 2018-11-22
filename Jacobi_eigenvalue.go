package main

import "math"

var iter int

func JacobiEigenvalue(A [][]float64) []float64 {
	var d, s, s1, s1t, temp [][]float64
	var theta float64
	flag := 1.
	n := len(A)
	s = EMatrix(n)

	for i := 0; i < len(A); i++ {
		d = append(d, []float64{})
		temp = append(temp, []float64{})
		for j := 0; j < n; j++ {
			d[i] = append(d[i], A[i][j])
			temp[i] = append(temp[i], 0.)
		}
	}
	for flag == 1 {
		flag = 0
		i, j := maxOffDiagonal(d)

		theta = 0.5 * math.Atan2(2*d[i][j], d[i][i]-d[j][j])

		s1 = givensRotation(theta, n, i, j)
		s1t = T(s1)

		temp = mul(s1t, d)

		d = mul(temp, s1)

		temp = mul(s, s1)

		s = mul(s, s1)

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					if math.Abs(d[i][j]) > EPS {
						flag = 1
					}
				}
			}
		}
		iter++
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
