package main

import (
	"./display"
	"./draw"
	"./matrix"
	"./parser"
)

func main() {
	screen := display.NewScreen()
	transform := matrix.NewMatrix()
	edges := [][]float64{
		{0.0},
		{0.0},
		{0.0},
		{1.0},
	}

	// defer display.DisplayScreen(screen)
	defer draw.DrawLines(edges, screen)

	parser.ParseFile("script", transform, edges, screen)

}
