// Package matrix provides functions for creating and manipulating matrices
package matrix

import (
	"fmt"
	"math"
	"strings"
)

func MakeIdentity(matrix [][]float64) {
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

func MultiplyMatrices(m1Ptr, m2Ptr *[][]float64) {
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

func ExtractColumn(matrix [][]float64, colIndex int) []float64 {
	col := make([]float64, len(matrix))

	for i, _ := range matrix {
		col[i] = matrix[i][colIndex]
	}

	return col
}

func dot(x, y []float64) float64 {
	output := 0.0
	for i, _ := range x {
		output += x[i] * y[i]
	}
	return output
}

func NewMatrix(params ...int) [][]float64 {
	rows := 4
	cols := 4

	if len(params) >= 2 {
		rows = params[0]
		cols = params[1]
	}

	matrix := make([][]float64, rows)
	for i, _ := range matrix {
		matrix[i] = make([]float64, cols)
	}

	return matrix
}

func PrintMatrix(matrix [][]float64) {
	output := ""

	for _, row := range matrix {
		for _, value := range row {
			output += fmt.Sprintf("%f%s", value, strings.Repeat(" ", 4 - len(fmt.Sprint(value))))
		}
		output += "\n"
	}

	fmt.Println(output)
}

func MakeTranslationMatrix(x, y, z float64) (m [][]float64) {
	m = NewMatrix()
	MakeIdentity(m)
	m[0][3], m[1][3], m[2][3] = x, y, z
	return
}

func MakeDilationMatrix(x, y, z float64) (m [][]float64) {
	m = NewMatrix()
	MakeIdentity(m)
	m[0][0], m[1][1], m[2][2] = x, y, z
	return
}

// MakeRotX creates a rotation matrix from a theta as the angle of rotation
// around the x axis. It returns the rotation matrix.
func MakeRotX(theta float64) (m [][]float64) {
	m = NewMatrix()
	radians := theta / 180 * math.Pi
	MakeIdentity(m)
	m[1][1], m[1][2] = math.Cos(radians), -math.Sin(radians)
	return
}
