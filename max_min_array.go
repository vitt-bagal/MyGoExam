package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	max, min := findMaxAndMinFromArray(arr)
	fmt.Println("Max n min number are:", max, min)
}

// find max and min from given array
func findMaxAndMinFromArray(arr [5]int) (int, int) {
	var max, min int = 0, 9999
	for i, _ := range arr {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}
