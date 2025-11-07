package main

import "fmt"

//print star pattern in traingle
func main() {
	n := 3
	for i := 0; i < 3; i++ {
		j := n
		for j > 0 {
			fmt.Printf(" ")
			j--
		}
		for k := 0; k <= i; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
		n--
	}
}
