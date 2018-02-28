// Contains display functions for screens

package main

const XRES = 500
const YRES = 500

func NewScreen() (screen [][][]int) {
	screen = make([][][]int, YRES)

	for i, _ := range screen {
		screen[i] = make([][]int, XRES)

		for j, _ := range screen[i] {
			screen[i][j] = make([]int, 3)
		}
	}

	return
}
