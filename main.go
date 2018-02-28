package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	m := NewMatrix()
	fmt.Println("\nBlank matrix =====")
	PrintMatrix(m)

	fmt.Println("Identity matrix =====")
	MakeIdentity(m)
	PrintMatrix(m)

	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	m2 := [][]int{
		{1, 2, 3, 4},
		{2, 6, 7, 8},
		{3, 10, 11, 12},
	}
	fmt.Println("Matrices to multiply:")
	PrintMatrix(m1)
	PrintMatrix(m2)

	fmt.Println("Product of multiplication:")
	MultiplyMatrices(&m1, &m2)
	PrintMatrix(m2)

	drawGopher()
}

func drawGopher() {
	screen := NewScreen()
	matrix := [][]int{
		{145},
		{79},
		{0},
		{1},
	}

	defer WriteScreenToPPM(screen)
	defer DrawLines(matrix, screen)

	AddEdge(matrix, 142, 82, 0, 162, 70, 0)

	// Yes, I had to build a HTML5 Canvas with an image of the Go Gopher drawn onto it
	// and use a bit of JS to determine these points.
	AddPoint(matrix, 189, 52, 0)
	AddPoint(matrix, 206, 39, 0)
	AddPoint(matrix, 230, 33, 0)
	AddPoint(matrix, 254, 33, 0)
	AddPoint(matrix, 272, 36, 0)
	AddPoint(matrix, 289, 45, 0)
	AddPoint(matrix, 298, 40, 0)
	AddPoint(matrix, 310, 35, 0)
	AddPoint(matrix, 334, 46, 0)
	AddPoint(matrix, 338, 61, 0)
	AddPoint(matrix, 332, 78, 0)
	AddPoint(matrix, 339, 90, 0)
	AddPoint(matrix, 353, 119, 0)
	AddPoint(matrix, 365, 150, 0)
	AddPoint(matrix, 377, 167, 0)
	AddPoint(matrix, 383, 187, 0)
	AddPoint(matrix, 400, 184, 0)
	AddPoint(matrix, 412, 190, 0)
	AddPoint(matrix, 406, 205, 0)
	AddPoint(matrix, 390, 202, 0)
	AddPoint(matrix, 398, 260, 0)
	AddPoint(matrix, 399, 294, 0)
	AddPoint(matrix, 393, 329, 0)
	AddPoint(matrix, 388, 362, 0)
	AddPoint(matrix, 400, 366, 0)
	AddPoint(matrix, 409, 379, 0)
	AddPoint(matrix, 398, 392, 0)
	AddPoint(matrix, 391, 382, 0)
	AddPoint(matrix, 378, 380, 0)
	AddPoint(matrix, 357, 399, 0)
	AddPoint(matrix, 330, 412, 0)
	AddPoint(matrix, 298, 427, 0)
	AddPoint(matrix, 263, 442, 0)
	AddPoint(matrix, 241, 443, 0)
	AddPoint(matrix, 226, 465, 0)
	AddPoint(matrix, 207, 464, 0)
	AddPoint(matrix, 214, 438, 0)
	AddPoint(matrix, 194, 424, 0)
	AddPoint(matrix, 185, 409, 0)
	AddPoint(matrix, 178, 414, 0)
	AddPoint(matrix, 172, 405, 0)
	AddPoint(matrix, 178, 400, 0)
	AddPoint(matrix, 173, 384, 0)
	AddPoint(matrix, 175, 336, 0)
	AddPoint(matrix, 168, 295, 0)
	AddPoint(matrix, 160, 266, 0)
	AddPoint(matrix, 143, 245, 0)
	AddPoint(matrix, 124, 211, 0)
	AddPoint(matrix, 117, 193, 0)
	AddPoint(matrix, 113, 180, 0)
	AddPoint(matrix, 88, 172, 0)
	AddPoint(matrix, 83, 156, 0)
	AddPoint(matrix, 86, 139, 0)
	AddPoint(matrix, 102, 129, 0)
	AddPoint(matrix, 112, 128, 0)
	AddPoint(matrix, 120, 105, 0)
}

func WriteScreenToPPM(screen [][][]int) {
	file, err := os.OpenFile("foo.ppm", os.O_CREATE | os.O_WRONLY, 0644)

	if (err != nil) {
		panic(err)
	}

	defer file.Close()

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("P3 %d %d 255\n", XRES, YRES))
	for i := 0; i < YRES; i++ {
		for j := 0; j < XRES; j++ {
			rgb := screen[i][j]
			buffer.WriteString(fmt.Sprintf("%d %d %d ", uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2])))
		}
	}

	file.WriteString(buffer.String())
}
