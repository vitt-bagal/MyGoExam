package main

import "fmt"

var stackSlice = []int{}

func main() {

	// Push element to stack
	push(1)
	push(2)
	push(3)
	push(4)
	fmt.Println("stack after 4 push operation: ", stackSlice)
	// pop element from stack
	pop()
	pop()
	fmt.Println("stack after 2 pop operation: ", stackSlice)
}

func pop() {
	l := len(stackSlice)
	stackSlice = stackSlice[:l-1]
}

func push(element int) {
	stackSlice = append(stackSlice, element)
}
