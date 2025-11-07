package main

import "fmt"

func main() {
	op := adder()
	for i := 0; i < 10; i++ {
		fmt.Println("Final val", op(i))
	}

}

func adder() func(y int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
