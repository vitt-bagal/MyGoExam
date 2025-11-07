package main

import "fmt"

// A number which leaves 1 as a result after a sequence of steps and in each step number is replaced by the sum of squares of its digit.
// For example, if we check whether 23 is a Happy Number, then sequence of steps are

// Step 1: 2×2+3×3 = 4+9 = 13 // Sum of square of each digit
// Step 2: 1×1+3×3 = 1+9 = 10
// Step 3: 1×1+0x0 = 1 (A Happy Number)

func main() {
	num := 2
	if checkHappyNum(num) {
		fmt.Printf("Num %d is happy number", num)
	} else {
		fmt.Printf("Num %d is not happy number", num)
	}
	fmt.Println(getDigits(num))
}

func checkHappyNum(num int) bool {
	for num != 1 {
		tmp := getDigits(num)
		fmt.Println(tmp)
		squareSum := 0
		for i := 0; i < len(tmp); i++ {
			squareSum = squareSum + tmp[i]*tmp[i]
		}
		num = squareSum
	}
	if num == 1 {
		return true
	} else {
		return false
	}
}

func getDigits(num int) []int {
	//var digits []int
	digits := []int{}
	for num != 0 {
		digits = append(digits, (num % 10))
		num = num / 10
	}
	return digits
}
