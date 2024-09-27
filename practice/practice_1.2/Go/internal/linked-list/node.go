package dll

import (
	"encoding/json"
	"fmt"
)

// List Node
//
// Has links to next/previous nodes and to the DoublyLinkedList
type Node[T comparable] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{Data: data}
}

func (node *Node[T]) HasNext() bool {
	return node.Next != nil
}

func (node *Node[T]) HasPrev() bool {
	return node.Prev != nil
}

// Stringer method
func (node *Node[T]) String() string {
	s, err := json.MarshalIndent(node.Data, "", "  ")
	if err != nil {
		println(err.Error())
	}
	str := "<node>\n"
	str = str + "data:" + string(s)
	if node.HasPrev() {
		prev, err := json.MarshalIndent(node.Prev.Data, "", "  ")
		if err != nil {
			println(err.Error())
		}
		prevString := string(prev)
		str = str + "\nprev:" + prevString
	}
	if node.HasNext() {
		next, err := json.MarshalIndent(node.Next.Data, "", "  ")
		if err != nil {
			println(err.Error())
		}
		nextString := string(next)
		str = str + "\nnext:" + nextString
	}
	str = str + "\n</node>"
	return str
}

func Testing() {
	fmt.Print("eeeeeeeeeeeeeeeeeeeeeeeee")
}
