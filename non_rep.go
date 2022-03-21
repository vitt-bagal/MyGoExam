/*Find the first non-repeating element in a given array arr of N integers.
Note: Array consists of only positive and negative integers and not zero.

Example 1:

Input : arr[] = {-1, 2, -1, 3, 2}
Output : 3
Explanation:
-1 and 2 are repeating whereas 3 is
the only number occuring once.
Hence, the output is 3.
Example 1:

Input : arr[] = {1, 1, 1}
Output : 0 */

package main

import "fmt"

func main() {
	var nonRepValue = 0
	arr := []int{-1, 2, -1, 3, 2}
	for i := 0; i < len(arr); i++ {
		if !isExist(arr, arr[i]) {
			nonRepValue = arr[i]
		}
	}
	fmt.Println(nonRepValue)
}

func isExist(a []int, ele int) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == ele {
			count++
		}
	}
	if count == 1 {
		return false
	}
	return true
}
