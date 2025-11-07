//Find duplicates in an array
//Input:
//N = 5
//a[] = {2,3,1,2,3}
//Output: 2 3 
//Explanation: 2 and 3 occur more than once
//in the given array.

package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 1, 2, 2, 3}
	fmt.Println("List of duplicate", findListDuplicates(s))
}

func findListDuplicates(s []int) []int {
	var duplicate []int
	m := make(map[int]int)
	for _, v := range s {
		k, ok := m[v]
		if ok {
			k++
			m[v] = k
			if m[v] == 2 {
				duplicate = append(duplicate, v)
			}
		} else {
			m[v] = 1
		}
	}
	if duplicate == nil {
		duplicate = append(duplicate, -1)
	}
	return duplicate
}
