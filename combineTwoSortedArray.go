package main

import "fmt"

func main() {
	// create sample input array a1 and a2
	// array a3 will be output array
	a1 := [4]int{1, 5, 6, 10}
	fmt.Println("Array a1 is", a1)
	a2 := [2]int{20, 234}
	fmt.Println("Array a2 is", a2)
	m := len(a1)
	n := len(a2)
	a3 := make([]int, (m + n))
	fmt.Println("Array a3 is before (a1,a2 merged)", a3)
	i, j, k := 0, 0, 0
	// Traverse both array simultanesouly
	for i < m && j < n {
		if a1[i] <= a2[j] {
			a3[k] = a1[i]
			i++
			k++
		} else {
			a3[k] = a2[j]
			j++
			k++
		}
	}

	// Copy all remaining element from a1 to a3
	for i < m {
		a3[k] = a1[i]
		i++
		k++
	}

	// Copy all remaining element from a2 to a3
	for j < n {
		a3[k] = a2[j]
		j++
		k++
	}
	fmt.Println("Comined merged array is", a3)
}
