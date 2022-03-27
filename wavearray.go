//Given a sorted array arr[] of distinct integers. Sort the array into a wave-like array and return it
//In other words, arrange the elements into a sequence such that arr[1] >= arr[2] <= arr[3] >= arr[4] <= arr[5].....
package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	waveArray(arr1)
	arr2 := []int{2, 4, 7, 8, 9, 10}
	waveArray(arr2)
}

func waveArray(n []int) {
	for i := 0; i < len(n)-1; {
		n[i], n[i+1] = n[i+1], n[i]
		i += 2
	}
	fmt.Println("Wave array is - ", n)
}
