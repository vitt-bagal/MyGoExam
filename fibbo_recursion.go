package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-2) + fibonacci(n-1)
	}
}

func main() {
	var num = fibonacci(10)
	fmt.Printf("\t %d", num)
}
