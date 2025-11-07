package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var xVisited, yVisited = false, false
	x, y, sum := 0, 1, 1
	return func() int {
		if x == 0 && !xVisited {
			xVisited = true
			return x
		} else if y == 1 && !yVisited {
			yVisited = true
			return y
		} else {
			sum = x + y
			x, y = y, sum
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("\t %d", f())
	}
}
