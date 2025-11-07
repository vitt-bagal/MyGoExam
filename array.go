package main

import "fmt"

func main() {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	fmt.Println("Array in go", arr)
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i] * 2
			count++
		}
	}
	for count < len(arr) {
		arr[count] = 0
		count++
	}
	fmt.Println("Array after chnage", arr)
}
