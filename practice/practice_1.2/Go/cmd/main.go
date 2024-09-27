package main

import (
	"fmt"

	"practice_1.2/internal/linked-list"
)

func main() {
	var node = dll.NewNode("1")
	node1 := dll.NewNode("2")
	var l = dll.NewDoublyLinkedList[string]()
	l.AppendLeft(node)
	l.AppendLeft(node1)
	fmt.Printf("%+v\n", l)
}
