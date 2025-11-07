package main

import (
	"fmt"
)

//https://www.geeksforgeeks.org/program-count-numbers-fingers/
func main() {
	a := 17
	fmt.Printf("Output of %v is %v \n", a, count_no_finger(a))

	b := 27
	fmt.Printf("Output of %v is %v \n", b, count_no_finger(b))
}

func count_no_finger(x int) int {
	r := x % 8
	if r == 1 || r == 5 {
		return r
	}
	if r == 0 || r == 2 {
		return 2
	}
	if r == 4 || r == 6 {
		return 4
	}
	if r == 3 || r == 7 {
		return 3
	}
	return x
}
