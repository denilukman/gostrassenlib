package gostrassenlib

import (
	"math"
)

func createMatrix(row_size int, col_size int) [][]int {
	M := make([][]int, row_size)
	for i := range M {
		M[i] = make([]int, col_size)
	}

	return M
}

func add(A [][]int, B [][]int) [][]int {
	row_size := len(A)
	col_size := len(A[0])

	C := createMatrix(row_size, col_size)

	for i := 0; i < row_size; i++ {
		for j := 0; j < col_size; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

func substract(A [][]int, B [][]int) [][]int {
	row_size := len(A)
	col_size := len(A[0])

	C := createMatrix(row_size, col_size)

	for i := 0; i < row_size; i++ {
		for j := 0; j < col_size; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}

func calculate(A [][]int, B [][]int) [][]int {
	C := createMatrix(2, 2)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			C[i][j] = 0
			for k := 0; k < 2; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

func powerOfTwo(n int) int {
	return int(math.Pow(2, math.Ceil(math.Log(float64(n))/math.Log(float64(2)))))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func strassen(A [][]int, B [][]int) [][]int {
	size := len(A)

	if size == 2 {
		C := calculate(A, B)
		return C
	}

	new_size := size / 2

	A11 := createMatrix(new_size, new_size)
	A12 := createMatrix(new_size, new_size)
	A21 := createMatrix(new_size, new_size)
	A22 := createMatrix(new_size, new_size)

	B11 := createMatrix(new_size, new_size)
	B12 := createMatrix(new_size, new_size)
	B21 := createMatrix(new_size, new_size)
	B22 := createMatrix(new_size, new_size)

	for i := 0; i < new_size; i++ {
		for j := 0; j < new_size; j++ {
			A11[i][j] = A[i][j]
			A12[i][j] = A[i][j+new_size]
			A21[i][j] = A[i+new_size][j]
			A22[i][j] = A[i+new_size][j+new_size]

			B11[i][j] = B[i][j]
			B12[i][j] = B[i][j+new_size]
			B21[i][j] = B[i+new_size][j]
			B22[i][j] = B[i+new_size][j+new_size]
		}
	}

	//M1 = (A11 + A22) * (B11 + B22)
	a := add(A11, A22)
	b := add(B11, B22)
	M1 := strassen(a, b)

	//M2 = (A21 + A22) * (B11)
	a = add(A21, A22)
	M2 := strassen(a, B11)

	//M3 = (A11) * (B12 - B11)
	b = substract(B12, B22)
	M3 := strassen(A11, b)

	//M4 = A22 * (B21 - B22)
	b = substract(B21, B11)
	M4 := strassen(A22, b)

	//M5 = (A11 + A12) * B22
	a = add(A11, A12)
	M5 := strassen(a, B22)

	//M6 = (A21 - A11) * (B11 + B12)
	a = substract(A21, A11)
	b = add(B11, B12)
	M6 := strassen(a, b)

	//M7 = (A12 - A22) * (B21 + B22)
	a = substract(A12, A22)
	b = add(B21, B22)
	M7 := strassen(a, b)

	//C11 = M1 + M4 - M5 + M7
	a = add(M1, M4)
	b = add(a, M7)
	C11 := substract(b, M5)

	//C12 = M3 + M5
	C12 := add(M3, M5)

	//C21 = M2 + M4
	C21 := add(M2, M4)

	//C22 = M1 - M2 + M3 + M6
	a = add(M1, M3)
	b = add(a, M6)
	C22 := substract(b, M2)

	C := createMatrix(size, size)

	for i := 0; i < new_size; i++ {
		for j := 0; j < new_size; j++ {
			C[i][j] = C11[i][j]
			C[i][j+new_size] = C12[i][j]
			C[i+new_size][j] = C21[i][j]
			C[i+new_size][j+new_size] = C22[i][j]
		}
	}

	return C
}

func Multiply(A [][]int, B [][]int) [][]int {
	row_size := len(A)
	col_size := len(B[0])

	n := powerOfTwo(max(max(len(A), len(A[0])), max(len(B), len(B[0]))))

	Astrassen := createMatrix(n, n)
	Bstrassen := createMatrix(n, n)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			Astrassen[i][j] = A[i][j]
		}
	}

	for i := 0; i < len(B); i++ {
		for j := 0; j < len(B[0]); j++ {
			Bstrassen[i][j] = B[i][j]
		}
	}

	Cstrassen := strassen(Astrassen, Bstrassen)

	//trim the 0s
	C := createMatrix(row_size, col_size)
	for i := 0; i < row_size; i++ {
		for j := 0; j < col_size; j++ {
			C[i][j] = Cstrassen[i][j]
		}
	}

	return C
}
