package main

import (
	"fmt"
	"math"
)

func gaussSolver(n int, A [][]float64, b []float64) []float64 {
	var i, j, k, l, m int
	x := make([]float64, n)
	//ETAPA DE ESCALONAMENTO
	for k = 0; k < n-1; k++ {
		max := math.Abs(A[k][k])
		indiceMax := k
		for i = k + 1; i < n; i++ {
			if max < math.Abs(A[i][k]) {
				max = math.Abs(A[i][k])
				indiceMax = i
			}
		}
		if indiceMax != k {

			for j = 0; j < n; j++ {
				temp := A[k][j]
				A[k][j] = A[indiceMax][j]
				A[indiceMax][j] = temp
			}
			temp := b[k]
			b[k] = b[indiceMax]
			b[indiceMax] = temp
		}
		if A[k][k] == 0 {
			fmt.Println("A matriz dos coeficientes é singular")
			return x
		} else {
			for m = k + 1; m < n; m++ {
				F := -A[m][k] / A[k][k]
				A[m][k] = 0
				b[m] = b[m] + F*b[k]

				for l = k + 1; l < n; l++ {
					A[m][l] = A[m][l] + F*A[k][l]
				}
			}
		}
	}
	for i = n - 1; i >= 0; i-- {
		x[i] = b[i]
		for j = i + 1; j < n; j++ {
			x[i] = x[i] - x[j]*A[i][j]
		}
		x[i] = x[i] / A[i][i]
	}
	return x
}

func main() {
	n := 3
	A := [][]float64{{2, 1, -1}, {1, 2, 1}, {1, 1, 1}}
	b := []float64{-3, 3, 2}
	var x []float64
	x = gaussSolver(n, A, b)
	fmt.Printf("x1 = %v\nx2 = %v\nx3 = %v\n", x[0], x[1], x[2])

	var A2 = [][]float64{{1, 1, 1}, {-2, 1, 1}, {1, 3, 1}}
	var b2 = []float64{2, 5, 4}
	x = gaussSolver(n, A2, b2)
	fmt.Printf("\nx1 = %v\nx2 = %v\nx3 = %v\n", x[0], x[1], x[2])

	A3 := [][]float64{{3, 2, -1}, {2, -2, 4}, {-1, 0.5, -1}}
	b3 := []float64{1, -2, 0}
	x = gaussSolver(n, A3, b3)
	fmt.Printf("\nx1 = %v\nx2 = %v\nx3 = %v\n", x[0], x[1], x[2])

	n = 2
	A4 := [][]float64{{2, 3}, {4, 9}}
	b4 := []float64{6, 15}
	x = gaussSolver(n, A4, b4)
	fmt.Printf("\nx1 = %v\nx2 = %v\n", x[0], x[1])

	n = 4
	A5 := [][]float64{{4, 1, 2, -3}, {-3, 3, -1, 4}, {-1, 2, 5, 1}, {5, 4, 3, -1}}
	b5 := []float64{-16, 20, -4, -10}
	x = gaussSolver(n, A5, b5)
	fmt.Printf("\nx1 = %v\nx2 = %v\nx3 = %v\nx4 = %v\n", x[0], x[1], x[2], x[3])

	n = 5
	A6 := [][]float64{{4, 1, 2, -3, 5}, {-3, 3, -1, 4, -2}, {-1, 2, 5, 1, 3}, {5, 4, 3, -1, 2}, {1, -2, 3, -4, 5}}
	b6 := []float64{-16, 20, -4, -10, 3}
	x = gaussSolver(n, A6, b6)
	fmt.Printf("\nx1 = %v\nx2 = %v\nx3 = %v\nx4 = %v\nx5 = %v\n", x[0], x[1], x[2], x[3], x[4])
}
