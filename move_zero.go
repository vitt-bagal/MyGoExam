// Move Zeroes Problem
//
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Example:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
package main

import "fmt"

func main() {
	ip := []int{0, 1, 0, 3, 12}
	//lastNonZeroFoundAt := 0
	for i := 0; i < len(ip); i++ {
		for j := i; j < len(ip); j++ {
			if ip[i] == 0 && ip[j] != 0 {
				ip[i], ip[j] = ip[j], ip[i]
			}
		}
	}
	fmt.Println(ip)
}
