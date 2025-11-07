package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Array in go", arr)
	fmt.Println("Array after reversing element", reverseArray(arr))

}

func reverseArray(a []int) []int {
	i := 0
	j := len(a) - 1
	for i < j {
		var temp int = a[i]
		a[i] = a[j]
		a[j] = temp
		i++
		j--
	}
	return a
}
