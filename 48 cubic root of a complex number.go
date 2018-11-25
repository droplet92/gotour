package main

import (
	"fmt"
	"math/cmplx"
)

// Cbrt returns cubic root of x
func Cbrt(x complex128) complex128 {
	z := complex128(1)
	for i := 0; i < 10; i++ {
		z = z - (z*z*z-x)/(3*z*z)
	}
	return z
}

func main() {
	fmt.Println(
		Cbrt(2),
		cmplx.Pow(Cbrt(2), complex128(3)),
	)
}
