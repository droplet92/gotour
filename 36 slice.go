package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]byte, dy)

	// init slice
	for i := range slice {
		subslice := make([]byte, dx)
	
		// init subslice
		for j := range subslice {
			subslice[j] = byte(i + j)
		}
		slice[i] = subslice
	}
	return slice
}

func main() {
    pic.Show(Pic)
}
