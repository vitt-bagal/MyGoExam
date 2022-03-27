package main

import "fmt"

// Need to solve using recursion
func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
	fmt.Printf("Len-%d and cap is-%d", len(nums), cap(nums))
}
func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	fmt.Println(nums)
	return len(nums) - count
}
