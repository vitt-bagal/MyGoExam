package main

//Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

//Example 1:
//Input: nums1 = [1,3], nums2 = [2]
//Output: 2.00000
//Explanation: merged array = [1,2,3] and median is 2.

//Example 2:
//Input: nums1 = [1,2], nums2 = [3,4]
//Output: 2.50000
//Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

//Example 3:
//Input: nums1 = [0,0], nums2 = [0,0]
//Output: 0.00000

//Example 4:
//Input: nums1 = [], nums2 = [1]
//Output: 1.00000

//Example 5:
//Input: nums1 = [2], nums2 = []
//Output: 2.00000

/*
//Constraints:

//nums1.length == m
//nums2.length == n
//0 <= m <= 1000
//0 <= n <= 1000
//1 <= m + n <= 2000
//-106 <= nums1[i], nums2[i] <= 106*/

import (
	"fmt"
)

func main() {
	n1 := []int{1, 3}
	n2 := []int{2}
	fmt.Println("Median of two array is", findMedianSortedArrays(n1, n2))
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Println("Median of two slice is", findMedianSortedSlice(s1, s2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var medOfArray float64
	//length := len(nums1) + len(nums1)
	//nums3 := [length]int{}
	m := len(nums1)
	n := len(nums2)
	nums3 := make([]int, m+n)

	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums3[k] = nums1[i]
			k++
			i++
		} else {
			nums3[k] = nums1[j]
			j++
		}
	}
	for i < len(nums1) {
		nums3[k] = nums1[i]
		k++
		i++
	}
	for j < len(nums2) {
		nums3[k] = nums1[j]
		k++
		j++

	}

	index := len(nums3) / 2
	if len(nums3)%2 == 0 {
		medOfArray = (float64)(nums3[index]+nums3[index+1]) / 2
	} else {
		medOfArray = (float64)(nums3[index])
	}
	fmt.Println("Final array is ", nums3)
	return medOfArray
}

func findMedianSortedSlice(nums1 []int, nums2 []int) float64 {
	var medOfArray float64
	//length := len(nums1) + len(nums1)
	//nums3 := [length]int{}
	m := len(nums1)
	n := len(nums2)
	nums3 := make([]int, m+n)

	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums3[k] = nums1[i]
			k++
			i++
		} else {
			nums3[k] = nums1[j]
			j++
			k++
		}
	}
	for i < len(nums1) {
		nums3[k] = nums1[i]
		k++
		i++
	}
	for j < len(nums2) {
		nums3[k] = nums1[j]
		k++
		j++
	}

	index := len(nums3) / 2
	if len(nums3)%2 == 0 {
		medOfArray = (float64)(nums3[index]+nums3[index+1]) / 2
	} else {
		medOfArray = (float64)(nums3[index])
	}
	fmt.Println("Final slice is ", nums3)
	return medOfArray
}
