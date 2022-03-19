// You can edit this code!
// Click here and start typing.

// Program to read value from channel if value is sent otherwise wait till 1 second and return
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go MyFunc(ch)
	ch <- 5
	time.Sleep(2 * time.Second)
}
func MyFunc(c chan int) {
	timer := time.NewTimer(time.Second)
	select {
	case val := <-c:
		fmt.Println("Value received...", val)
	case <-timer.C:
		fmt.Println("No value received from channel in 1 second ...")
	}
}
