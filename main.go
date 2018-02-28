package main

import (
	"bytes"
	"fmt"
	"os"

	"./display"
	"./draw"
	"./matrix"
)

func main() {
	// defer WriteScreenToPPM(screen)
	// defer draw.DrawLines(matrix, screen)

	screen = display.NewScreen()
	edges, transform := matrix.NewMatrix(), matrix.NewMatrix()
}

func WriteScreenToPPM(screen [][][]int) {
	file, err := os.OpenFile("foo.ppm", os.O_CREATE | os.O_WRONLY, 0644)

	if (err != nil) {
		panic(err)
	}

	defer file.Close()

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("P3 %d %d 255\n", display.XRES, display.YRES))
	for i := 0; i < display.YRES; i++ {
		for j := 0; j < display.
			XRES; j++ {
			rgb := screen[i][j]
			buffer.WriteString(fmt.Sprintf("%d %d %d ", uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2])))
		}
	}

	file.WriteString(buffer.String())
}
