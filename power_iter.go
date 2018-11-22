package main

import (
	"fmt"
	"math"
)

const EPS = 0.000001

func lMax(A [][]float64) float64 {
	var pr, nt, m []float64
	tempM := 0.
	for i := 0; i < len(A); i++ {
		pr = append(pr, 1.)
		nt = append(nt, 0.)
		m = append(m, 1.)
	}
	for math.Abs(m[0]-tempM) >= EPS {
		tempM = m[0]
		nt = mulMatVec(A, pr)
		m[0] = nt[0] / pr[0]
		pr = nt
	}
	return m[0]
}

func lMin(A [][]float64) float64 {
	n := len(A)
	E := mulMatSc(EMatrix(n), lMax(A))
	minusMat(E, A)
	l := lMax(A) - lMax(E)
	//l := lMax(A) - normMatrix(A)
	return l
}

func lminmod(A [][]float64) float64 {
	var lminmod float64
	n := len(A)
	max := lMax(A)
	C := minusMat(EMatrix(n), MatDivSc(mul(A, A), math.Pow(max, 2)))
	lminmod = math.Sqrt(math.Pow(max, 2) * (1 - lMax(C)))
	return math.Sqrt(lminmod)
}

func main() {
	//A := MyMatrix(3)
	C := symmetricM(5)
	printMatrix(C)
	//A := [][]float64{{1, 2}, {2, 1}}
	//fmt.Printf("%v\n", lMax(A))
	//fmt.Printf("%v\n", lMin(A))
	//JacobiEigenvalue(A)
	//B := [][]float64{{1, 2, 3}, {4, 6, 7}, {5, 8, 9}}
	//T(B)
	fmt.Printf("%v\n", lMax(C))
	fmt.Printf("%v\n", lMin(C))
	fmt.Printf("%v\n", lminmod(C))
	JacobiEigenvalue(C)
	//fmt.Print("\n", iter)
}
