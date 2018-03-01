package main

import (
	"./display"
	"./draw"
)

func main() {
	screen := display.NewScreen()
	edges := [][]float64{
		{0.0},
		{0.0},
		{0.0},
		{1.0},
	}

	defer display.DisplayScreen(screen)
	defer draw.DrawLines(edges, screen)

	draw.AddPoint(edges, 100, 0, 0)
	draw.AddPoint(edges, 100, 100, 0)
	draw.AddPoint(edges, 0, 100, 0)
}
