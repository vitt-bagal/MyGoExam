// Program to  print array elemt with sum of digit is more

package main

import "fmt"

func main() {
	arr := [6]int{45, 66, 78, 99, 111, 02}
	fmt.Println("Element having max digit sum is ", findMaxDigit(arr))
}

func findMaxDigit(a [6]int) int {
	maxSum := 0
	max := 0
	for i := 0; i < len(a); i++ {
		val := sumOfDigit(a[i])
		if val > maxSum {
			maxSum = val
			max = a[i]
		} else {
			continue
		}
	}
	return max
}

func sumOfDigit(d int) int {
	sum := 0
	for d > 1 {
		d = d % 10
		sum = sum + d
		d = d / 10
	}
	return sum
}
