package main

// https://go.dev/tour/flowcontrol/8

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var lo, hi, z float64 = 1, 1<<32 - 1, -1
	for lo <= hi {
		var mid float64 = (lo + hi) / 2
		if mid*mid <= x {
			z = mid
			lo = mid + 0.00000001
		} else {
			hi = mid - 0.00000001
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
