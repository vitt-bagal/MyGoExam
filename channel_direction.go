// Write a function that accepts two channels as input,
// receives a value from one channel and then sends that value to the other channel
package main

import "fmt"

func main() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c1 <- 4
	sendAndReceive(c1, c2)
	fmt.Println(<-c2)
}

func sendAndReceive(s <-chan int, r chan<- int) {
	n := <-s
	r <- n
}
