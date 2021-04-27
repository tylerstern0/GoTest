package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("pixel select version 0.2")
}

// PixelSelect accepts the dimensions of the rectangle from which to select pixels &
// PixelSelect prints a random sequence of every point in that rectangle
func PixelSelect(rc_len, rc_wid int) {
	// Input format check
	if rc_len <= 0 || rc_wid <= 0 {
		panic(errors.New("Rectangle dimensions must be positive numbers."))
	}

	// Create an array of vertice names

	// Print each vertex array element in the order of the shuffled index

	return
}
