package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter number to reverse: ")
	fmt.Scanf("%d", &num)
	fmt.Println("Reverse of entered number:", reverse(num))

}
func reverse(x int) int {
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	return rev
}
