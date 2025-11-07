package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hi in main")
	go hello("there")
	fmt.Println("Again called")
	time.Sleep(10 * time.Millisecond)
}

func hello(s string) {
	fmt.Printf("Hi in goroutine- %v", s)

}
