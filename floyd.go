// program to print floyd triangle
//1
//2 3
//4 5 6
//7 8 9 10

package main

import "fmt"

func main() {
	var temp = 0
	fmt.Println("Floyd Triangle")
	for i := 0; i < 4; i++ {
		for j := 0; j <= i; j++ {
			temp++
			fmt.Print(temp)
			fmt.Printf("\t")
		}
		fmt.Printf("\n")
	}
}
