// Package draw contains useful functions to manipulate the edge matrix and draw
// it onto a screen.
package draw

import (
	"../display"
	"../matrix"
)

var DefaultDrawColor []int = []int{0, 0, 0}

// DrawLines draws an edge matrix onto a screen.
func DrawLines(edges [][]float64, screen [][][]int) {
	for i := 0; i < len(edges[0]) - 1; i++ {
		point := matrix.ExtractColumn(edges, i)
		nextPoint := matrix.ExtractColumn(edges, i + 1)
		x0, y0 := point[0], point[1]
		x1, y1 := nextPoint[0], nextPoint[1]
		DrawLine(screen, x0, y0, x1, y1)
	}
}

// AddPoint adds a point to an edge matrix.
func AddPoint(m [][]float64, x, y, z float64) {
	m[0] = append(m[0], x)
	m[1] = append(m[1], y)
	m[2] = append(m[2], z)
	m[3] = append(m[3], 1)
}

// AddEdge adds an edge (two points) to an edge matrix.
func AddEdge(m [][]float64, params ...float64) {
	x0, y0, z0 := params[0], params[1], params[2]
	x1, y1, z1 := params[3], params[4], params[5]
	AddPoint(m, x0, y0, z0)
	AddPoint(m, x1, y1, z1)
}

// DrawLine draws a line from (x0, y0) to (x1, y1) onto a screen.
func DrawLine(screen [][][]int, x0, y0, x1, y1 float64) {
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

	slope := A / (-B)
	var d float64

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

// plot draws a point (x, y) onto a screen with the default draw color.
func plot(screen [][][]int, x, y float64) {
	newX, newY := float64ToInt(x), display.YRES - float64ToInt(y) - 1
	if (newX >= 0 && newX < display.XRES && newY >= 0 && newY < display.YRES) {
		screen[newY][newX] = DefaultDrawColor[:]
	}
}

// DrawLineFromParams gets arguments from a params slice.
func DrawLineFromParams(screen [][][]int, params ...float64) {
	if (len(params) >= 4) {
		DrawLine(screen, params[0], params[1], params[2], params[3])
	}
}

// float64ToInt rounds a float64 without truncating it. It returns an int.
func float64ToInt(f float64) int {
	if (f - float64(int(f)) < 0.5) {
		return int(f)
	}
	return int(f + 1)
}
