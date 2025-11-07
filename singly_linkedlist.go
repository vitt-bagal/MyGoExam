package main

import "fmt"

//TODO: Write multiple function to insert,delete
type Node struct {
	val  int
	next *Node
}

func main() {
	fmt.Println("Single linked list program")
	HEAD := Node{}
	//fmt.Println("Head node", HEAD)
	SecNode := Node{}
	ThirdNode := Node{}
	HEAD.val = 1
	HEAD.next = &SecNode
	SecNode.val = 2
	SecNode.next = &ThirdNode
	ThirdNode.val = 3
	ThirdNode.next = nil
	printList(&HEAD)
}

func printList(n *Node) {
	for n != nil {
		fmt.Println(n.val)
		n = n.next
	}
}
