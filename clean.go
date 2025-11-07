package main

import "fmt"

type cleaner interface {
	clean()
}

func clean() {
	fmt.Println("Inside clean")
}
func main() {
	fmt.Println("Inside main")

	c.clean()
}
