package main

import (
	"fmt"
	"time"
)

//type error Error

func main() {
	c := make(chan error)
	go runner(c)
	fmt.Println(<-c)
}

func runner(c chan error) {
	time.Sleep(time.Second)
	err := fmt.Errorf("Error in runner..")
	c <- err
}
