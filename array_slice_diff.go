package main

import (
	"fmt"
)

func main() {

	a1 := [3]int{1, 2, 3}
	fmt.Println("array1 before function call", a1)
	fmt.Println("array1 after  function call", arrayVal(a1))

	a2 := []int{1, 2, 3}
	fmt.Println("slice before function call", a2)
	fmt.Println("slice after function call", sliceVal(a2))
}

func arrayVal(y [3]int) [3]int {
	y[2] = y[2] * 2
	return y
}

func sliceVal(x []int) []int {
	x[2] = x[2] * 2
	return x
}
