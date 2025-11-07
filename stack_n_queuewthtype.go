package main

import (
	"fmt"
	"os"
)

type stack []int

type queue []int

func main() {
	var s stack
	fmt.Println("stack before push operation:", s)
	// push to stack
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println("stack after push operation:", s)
	//remove top element from stack
	s.pop()
	fmt.Println("stack after pop operation:", s)

	var q queue
	fmt.Println("queue before enq operation:", q)
	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	fmt.Println("queue after enq operation:", q)
	q.dQueue()
	fmt.Println("queue after dq operation:", q)
}

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() {
	if len(*s) == 0 {
		fmt.Println("stack is empty")
		os.Exit(1)
	}
	*s = (*s)[:len(*s)-1]
}

func (q *queue) enQueue(n int) {
	*q = append(*q, n)
}

func (q *queue) dQueue() {
	if len(*q) == 0 {
		fmt.Println("queue is empty")
		os.Exit(1)
	}
	*q = (*q)[1:]
}
