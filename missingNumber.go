package main

import (
	"fmt"
)

func main() {
	arr := []int{6, 1, 2, 8, 3, 4, 7, 10, 5}
	expSum, actualSum := 0, 0
	for i := 0; i < len(arr); i++ {
		actualSum += arr[i]
	}
	n := len(arr) + 1
	expSum = n * (n + 1) / 2
	fmt.Println("actual sum", actualSum)
	fmt.Println("expSum", expSum)
	fmt.Println("Misssing element", (expSum - actualSum))
}
