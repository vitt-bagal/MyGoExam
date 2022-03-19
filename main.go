// Program to print sum of first 1000 number with the help of 5 goroutines
// This program also solves race condition using mutex

package main

import (
	"fmt"
	"sync"
)

// declare initial value
var start, end, sum = 0, 200, 0

// goroutine function
func process(i int, wg *sync.WaitGroup, mut *sync.Mutex) {
	fmt.Println("started Goroutine ", i)
	mut.Lock() //  lock the values
	for j := start; j < end; j++ {
		sum += j
	}
	start, end = end, end+200
	mut.Unlock() // unlock the values
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 5 // 5 goroutine
	var wg sync.WaitGroup
	var mut sync.Mutex
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg, &mut)
	}
	wg.Wait()
	fmt.Println("Final sum:-", sum)
	fmt.Println("All go routines finished executing")
}
