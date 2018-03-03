// Package matrix provides functions for creating and manipulating matrices.
package matrix

import (
	"fmt"
	"math"
	"strings"
)

// MakeIdentity makes a matrix an identity matrix.
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

// MultiplyMatrices multiples two matrices and stores it in the second matrix
// given.
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

// ExtractColumn extracts the column of a matrix. It returns that column as
// a slice.
func ExtractColumn(matrix [][]float64, colIndex int) []float64 {
	col := make([]float64, len(matrix))

	for i, _ := range matrix {
		col[i] = matrix[i][colIndex]
	}

	return col
}

// dot receives two slices as vectors. It returns their dot product.
func dot(x, y []float64) float64 {
	output := 0.0
	for i, _ := range x {
		output += x[i] * y[i]
	}
	return output
}

// NewMatrix creates a new float64 matrix. The default row and column size is 4.
// It returns the new matrix.
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

// PrintMatrix prints a float64 matrix.
func PrintMatrix(matrix [][]float64) {
	output := ""

	for _, row := range matrix {
		for _, value := range row {
			floatString := fmt.Sprintf("%.2f", value)
			spaces := strings.Repeat(" ", 8 - len(floatString))
			output += floatString + spaces
		}
		output += "\n"
	}

	fmt.Println(output)
}

// MakeTranslationMatrix creates a translation matrix using x, y, and z as the
// translation offsets.
func MakeTranslationMatrix(params ...float64) (m [][]float64) {
	m = NewMatrix()
	MakeIdentity(m)
	m[0][3], m[1][3], m[2][3] = params[0], params[1], params[2]
	return
}

// MakeDilationMatrix creates a dilation matrix using x, y, and z as the
// dilation offsets.
func MakeDilationMatrix(params ...float64) (m [][]float64) {
	m = NewMatrix()
	MakeIdentity(m)
	m[0][0], m[1][1], m[2][2] = params[0], params[1], params[2]
	return
}

// MakeRotX creates a rotation matrix using theta as the angle of rotation and
// X as the axis of rotation. It returns the rotation matrix.
func MakeRotX(theta float64) (m [][]float64) {
	m = NewMatrix()
  radians := theta / 180 * math.Pi
	sin, cos := math.Sin(radians), math.Cos(radians)
	MakeIdentity(m)
	m[1][1], m[1][2] = cos, -sin
	m[2][1], m[2][2] = sin, cos
	return
}

// MakeRotY creates a rotation matrix using theta as the angle of rotation and
// Y as the axis of rotation. It returns the rotation matrix.
func MakeRotY(theta float64) (m [][]float64) {
	m = NewMatrix()
  radians := theta / 180 * math.Pi
	sin, cos := math.Sin(radians), math.Cos(radians)
	MakeIdentity(m)
	m[0][0], m[0][2] = cos, sin
	m[2][0], m[2][2] = -sin, cos
	return
}

// MakeRotZ creates a rotation matrix using theta as the angle of rotation and
// Z as the axis of rotation. It returns the rotation matrix.
func MakeRotZ(theta float64) (m [][]float64) {
	m = NewMatrix()
  radians := theta / 180 * math.Pi
	sin, cos := math.Sin(radians), math.Cos(radians)
	MakeIdentity(m)
	m[0][0], m[0][1] = cos, -sin
	m[1][0], m[1][1] = sin, cos
	return
}
