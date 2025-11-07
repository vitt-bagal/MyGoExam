package main

import (
	"fmt"
)

//https://practice.geeksforgeeks.org/problems/trapping-rain-water-1587115621/1
func main() {
	//arr := []int{3, 0, 0, 2, 0, 4} // o/p - 10
	//arr := []int{6, 9, 9} // op - 0
	//arr := []int{7, 4, 0, 9} // op - 10
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("Trapper water is", trapWater(arr))
}

func trapWater(num []int) int {
	max, secMax, trappedwater := 0, 0, 0
	for i := 0; i < len(num); i++ {
		if num[i] > max {
			if secMax < max {
				secMax = max
			}
			max = num[i]
		}
	}
	//fmt.Printf("First max- %d n Sec max- %d\n", max, secMax)
	for i := 0; i < len(num); i++ {
		//	min := math.Min(float64(secMax), float64(num[i]))
		if num[i] <= secMax {
			trappedwater += (secMax - num[i] - int(min))

		}
	}
	return trappedwater
}
