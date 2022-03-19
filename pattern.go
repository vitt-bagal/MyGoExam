package main

import "fmt"

func main() {
	i, k := 0, 1
	n := 5
	for i < 5 {
		// Loop for handling spaces
		for m := 0; m < n; m++ {
			fmt.Print(" ")
		}
		if n == 5 {
			fmt.Print(" ")
		}
		j := 1
		// Main loop to print *
		for j <= k {
			fmt.Print("* ")
			j++
		}
		k += 2
		fmt.Println()
		i++
		n--
	}
}
