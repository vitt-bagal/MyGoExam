package main

import (
	"errors"
	"fmt"
)

//type error Error

func main() {
	c := make(chan error)
	s1 := "abc"
	var p = &s1
	go runner(p, c)
	go runner(nil, c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func runner(name *string, c chan error) {
	//time.Sleep(time.Second) // not needed as default send n receive on unbuffered channel is blocking operation
	if name == nil {
		//err := fmt.Errorf("name shd no be nil:", name)
		err := errors.New("name shd no be nil")
		c <- err
	}
	c <- nil
}
