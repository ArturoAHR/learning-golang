package main

// https://go.dev/tour/moretypes/18

import "golang.org/x/tour/pic"

func calculate(x, y int) uint8 {
	result := x
	for i := 0; i < y; i++ {
		result = x*x + y*i
	}
	return uint8(result)
}

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		image[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			image[i][j] = calculate(i, j)
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
