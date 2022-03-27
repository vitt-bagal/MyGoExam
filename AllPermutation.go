package main

// Program to take input array with integer 4 element and
// Print all posible values of valid time entries by assuming first two entries as hr and last as min
// hr -> 00 - 23
// min  -> 00 - 59
import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Inside permutation")
	var num = []int{1, 2, 3, 4}
	permute(0, len(num)-1, num)
	//	for i, v := range numSlice {
	//fmt.Printf("row %d with value - %v\n", i, v)
	//	}

}

func permute(s, e int, n []int) {
	if s == e {
		//fmt.Println(n) // Print the permutated values
		hr := strconv.Itoa(n[0]) + strconv.Itoa(n[1])  // convert first two entries as hour from num
		min := strconv.Itoa(n[2]) + strconv.Itoa(n[3]) // convert last two entries as hour from num
		if hr >= "00" && hr < "24" && min >= "00" && min < "60" {
			tf := hr + ":" + min
			fmt.Println("valid time format - ", tf)
		}
	} else {
		for i := s; i <= e; i++ {
			n[s], n[i] = n[i], n[s]
			permute(s+1, e, n)
			n[s], n[i] = n[i], n[s]
		}
	}
}
