package main

import "fmt"

func main() {
	nums1 := 5
	nums2 := 5
	sqaureOf2 := func(n1, n2 int) int {
		return nums1 * nums2
	}(nums1, nums2)

	fmt.Println(sqaureOf2)
}
