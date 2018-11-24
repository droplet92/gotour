package main

import (
	"fmt"
	"math"
)

const initval, delta = 1.0, 0.00000001

func Sqrt(x float64) float64 {
	var z, before float64 = initval, 0
	cnt := 0

	for math.Abs(z - before) > delta {
		before = z
		// f(z)=z^2, f'(z)=2*z
		z = z - (z*z - x) / (2*z)
		cnt++
	}
	fmt.Printf("sqrt(%g) loops %d times\n", x, cnt)
	return z
}

func main() {
	fmt.Println(
		math.Sqrt(2) - Sqrt(2),
		math.Sqrt(3) - Sqrt(3),
		math.Sqrt(4) - Sqrt(4),
		math.Sqrt(10) - Sqrt(10),
		math.Sqrt(37) - Sqrt(37),
	)
}