package main

import "fmt"

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.

//Example 1:

//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Output: Because nums[0] + nums[1] == 9, we return [0, 1].

func main() {
	arr := []int{3, 3}
	retArr := twoSum(arr, 6)
	fmt.Println(retArr)
}

func twoSum(nums []int, target int) []int {
	retA := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				retA = append(retA, i, j)
				break
			}
		}
	}
	return retA
}
