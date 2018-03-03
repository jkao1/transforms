package main

import (
	"./display"
	"./parser"
)

func main() {
	screen := display.NewScreen()
	transform := make([][]float64, 0)
	edges := make([][]float64, 4)

	parser.ParseFile("art", transform, edges, screen)

}
