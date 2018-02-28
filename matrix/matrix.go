package matrix

import (
	"fmt"
	"strings"
)

func MakeIdentity(matrix [][]int) {
	for i, row := range matrix {
		for j, _ := range row {
			if i == j {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}
		}
	}
}

func MultiplyMatrices(m1Ptr, m2Ptr *[][]int) {
	m1, m2 := *m1Ptr, *m2Ptr
	product := NewMatrix(len(m1), len(m2[0]))

	for i, row := range m1 {
		for j := 0; j < len(m2[0]); j++ {
			col := ExtractColumn(m2, j)
			product[i][j] = dot(row, col)
		}
	}

	*m2Ptr = product
}

func ExtractColumn(matrix [][]int, colIndex int) []int {
	col := make([]int, len(matrix))

	for i, _ := range matrix {
		col[i] = matrix[i][colIndex]
	}

	return col
}

func dot(x, y []int) int {
	output := 0
	for i, _ := range x {
		output += x[i] * y[i]
	}
	return output
}

func NewMatrix(params ...int) [][]int {
	rows := 4
	cols := 4

	if len(params) >= 2 {
		rows = params[0]
		cols = params[1]
	}

	matrix := make([][]int, rows)
	for i, _ := range matrix {
		matrix[i] = make([]int, cols)
	}

	return matrix
}

func PrintMatrix(matrix [][]int) {
	output := ""

	for _, row := range matrix {
		for _, value := range row {
			output += fmt.Sprintf("%d%s", value, strings.Repeat(" ", 4 - len(fmt.Sprint(value))))
		}
		output += "\n"
	}

	fmt.Println(output)
}

func MakeTranslationMatrix(x, y, z int) (m [][]int) {
	m = NewMatrix() // 4x4
	MakeIdentity(m)
	m[0][3], m[1][3], m[2][3] = x, y, z
	return m
}

func MakeDilationMatrix(x, y, z int) (m [][]int) {
	m = NewMatrix() // 4x4
	MakeIdentity(m)
	m[0][0], m[1][1], m[2][2] = x, y, z
	return m
}

// TODO: rotation matrix
