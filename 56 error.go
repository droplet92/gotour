package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

const initval, delta = 1.0, 0.00000001

func Sqrt(f float64) (float64, error) {
	// if f is a negative number,
	// then returns ErrNegativeSqrt
	if f < 0 {
		return f, ErrNegativeSqrt(f)
	}

	var z, before float64 = initval, 0
	cnt := 0

	for math.Abs(z-before) > delta {
		before = z
		// f(z)=z^2, f'(z)=2*z
		z = z - (z*z-f)/(2*z)
		cnt++
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
