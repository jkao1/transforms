package main

var DEFAULT_DRAW_COLOR []int = []int{255, 255, 255}

func DrawLines(matrix [][]int, screen [][][]int) {
	for i := 0; i < len(matrix[0]) - 1; i++ {
		point := ExtractColumn(matrix, i)
		nextPoint := ExtractColumn(matrix, i + 1)
		x0, y0 := point[0], point[1]
		x1, y1 := nextPoint[0], nextPoint[1]
		DrawLine(screen, x0, y0, x1, y1)
	}
	firstPoint := ExtractColumn(matrix, 0)
	lastPoint := ExtractColumn(matrix, len(matrix[0]) - 1)
	DrawLine(screen, firstPoint[0], firstPoint[1], lastPoint[0], lastPoint[1])
}

func AddPoint(matrix [][]int, x, y, z int) {
	matrix[0] = append(matrix[0], x)
	matrix[1] = append(matrix[1], y)
	matrix[2] = append(matrix[2], z)
	matrix[3] = append(matrix[3], 1)
}

func AddEdge(matrix [][]int, x0, y0, z0, x1, y1, z1 int) {
	AddPoint(matrix, x0, y0, z0)
	AddPoint(matrix, x1, y1, z1)
}

func DrawLine(screen [][][]int, x0, y0, x1, y1 int) {
	if x1 < x0 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	A := y1 - y0
	B := x0 - x1
	x := x0
	y := y0

	if B == 0 { // vertical line
		if y1 < y0 {
			y0, y1 = y1, y0
		}

		y = y0
		for y <= y1 {
			plot(screen, x, y)
			y++
		}

		return
	}

	slope := float64(A) / float64(-B)
	var d int

	if slope >= 0 && slope <= 1 { // octant 1
		d = 2*A + B
		for x <= x1 && y <= y1 {
			plot(screen, x, y)
			if d > 0 {
				y++
				d += 2*B
			}
			x++
			d += 2*A
		}
	}

	if slope > 1 { // octant 2
		d = A + 2*B
		for x <= x1 && y <= y1 {
			plot(screen, x, y)
			if d < 0 {
				x++
				d += 2*A
			}
			y++
			d += 2*B
		}
	}

	if slope < 0 && slope >= -1 { // octant 8
		d = 2*A - B
    for x <= x1 && y >= y1 {
			plot(screen, x, y)
			if d < 0 {
				y--
				d -= 2*B
			}
			x++
			d += 2*A
		}
	}

	if slope < -1 { // octant 7
		d = A - 2*B
		for x <= x1 && y >= y1 {
			plot(screen, x, y)
			if d > 0 {
				x++
				d += 2*A
			}
			y--
			d -= 2*B
		}
	}
}

func plot(screen [][][]int, x, y int) {
	if x >= 0 && x < XRES && y >= 0 && y < YRES {
		screen[y][x] = DEFAULT_DRAW_COLOR[:]
	}
}
