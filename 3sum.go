// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.
// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
package main

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeSum(nums)
}
func threeSum(nums []int) [][]int {
	var tsum = [][]int{}
	//tsum := make([][]int, 3)
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
					tsum[count] = make([]int, 3)
					//tsum[count] = append(tsum[count], nums[i])
					//tsum[count] = append(tsum[count], nums[j])
					//tsum[count] = append(tsum[count], nums[k])
					tsum[count][0] = nums[i]
					tsum[count][1] = nums[j]
					tsum[count][2] = nums[k]
					count++
				}
			}
		}
	}
	return tsum
}
