package main

// https://go.dev/tour/methods/20

import (
	"fmt"
)

type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot process square root of: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrorNegativeSqrt(x)
	}
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
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
