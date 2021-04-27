package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("pixel select version 0.2")
	PixelSelect(10, 5)
}

// PixelSelect accepts the dimensions of the rectangle from which to select pixels &
// PixelSelect prints a random sequence of every point in that rectangle
func PixelSelect(rc_wid, rc_hei int) {
	// Input format check
	if rc_hei <= 0 || rc_wid <= 0 {
		panic(errors.New("Rectangle dimensions must be positive numbers."))
	}

	seq_len := rc_hei * rc_wid
	//fmt.Println("expected number of elements: " + strconv.Itoa(seq_len))

	// Create an array of vertex names
	v_names := make([]string, seq_len)
	v_ind := 0
	for W := 0; W < rc_wid; W++ {
		for H := 0; H < rc_hei; H++ {
			v_names[v_ind] = strconv.Itoa(W) + ", " + strconv.Itoa(H)
			v_ind += 1
		}
	}

	//fmt.Println("v_names length: " + strconv.Itoa(len(v_names)))
	//fmt.Println("First element: " + v_names[0])

	// Create index
	sh_ind := make([]int, seq_len)
	for a := 0; a < seq_len; a++ {
		sh_ind[a] = a
	}

	// Shuffle index

	// Print each vertex array element in the order of the shuffled index
	// NB: currently prints sh_ind elements
	for pr_ind := 0; pr_ind < seq_len; pr_ind++ {
		fmt.Println(sh_ind[pr_ind])
	}

	return
}
