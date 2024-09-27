package dll

import (
	"errors"
)

// Doubly linked list structure
type DoublyLinkedList[T comparable] struct {
	First_node *Node[T] // First item of the dll that has the 0 index
	Current    *Node[T] // Current item of the dll
	Len        uint     // Length of the dll
}

// Doubly linked list constructor
func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		Len: 0,
	}
}

// Set current Node to next node if possible and return it
func (l *DoublyLinkedList[T]) Next() *Node[T] {
	if l.Len == 0 {
		return nil
	}
	if l.Current.HasNext() {
		l.Current = l.Current.Next
	}
	return l.Current
}

func (l *DoublyLinkedList[T]) HasNext() bool {
	return l.Current.HasNext()
}

// Set current Node to previous node if possible and return it
func (l *DoublyLinkedList[T]) Prev() *Node[T] {
	if l.Len == 0 {
		return nil
	}
	if l.Current.HasPrev() {
		l.Current = l.Current.Prev
	}
	return l.Current
}

func (l *DoublyLinkedList[T]) HasPrev() bool {
	return l.Current.HasPrev()
}

// Get last Node
func (l *DoublyLinkedList[T]) Last() *Node[T] {
	if l.Len == 0 {
		return nil
	}
	var temp *Node[T] = l.First_node
	for temp.HasNext() {
		temp = temp.Next
	}
	return temp
}

// Append to the left of the dll
//
// Receives ether T type or *Node[T]
func (l *DoublyLinkedList[T]) AppendLeft(dataOrNode interface{}) {
	switch dataOrNode := dataOrNode.(type) {
	case *Node[T]:
		if l.Len == 0 {
			l.First_node = dataOrNode // [node]<0>
			l.Current = dataOrNode    // ^[node]<0>
			l.Len++
			return
		}
		dataOrNode.Next = l.First_node // [node]     -> [first_node]<0>
		l.First_node.Prev = dataOrNode // [node]    <-> [first_node]<0>
		l.First_node = dataOrNode      // [node]<0> <-> [first_node]<1>
		l.Len++                        // increase the length of the dll
	case T:
		node := NewNode(dataOrNode)
		if l.Len == 0 {
			l.First_node = node // [node]<0>
			l.Current = node    // ^[node]<0>
			l.Len++
			return
		}
		node.Next = l.First_node // [node]     -> [first_node]<0>
		l.First_node.Prev = node // [node]    <-> [first_node]<0>
		l.First_node = node      // [node]<0> <-> [first_node]<1>
		l.Len++                  // increase the length of the dll
	default:
		panic("Wrong type provided")
	}

}

// Append to the right of the dll
//
// Receives ether T type or *Node[T]
func (l *DoublyLinkedList[T]) AppendRight(dataOrNode interface{}) {
	switch dataOrNode := dataOrNode.(type) {
	case *Node[T]:
		if l.Len == 0 {
			l.First_node = dataOrNode // [node]<0>
			l.Current = dataOrNode    // ^[node]<0>
			l.Len++
			return
		}
		temp := l.Last()
		if temp != nil {
			temp.Next = dataOrNode // [last]  -> [node]
			dataOrNode.Prev = temp // [last] <-> [node]
			l.Len++
		}
	case T:
		node := NewNode(dataOrNode) // create a new Node with data
		if l.Len == 0 {
			l.First_node = node // [node]<0>
			l.Current = node    // ^[node]<0>
			l.Len++
			return
		}
		temp := l.Last()
		if temp != nil {
			temp.Next = node // [last]  -> [node]
			node.Prev = temp // [last] <-> [node]
			l.Len++
		}
	default:
		panic("Wrong type provided")
	}

}

func (l *DoublyLinkedList[T]) Insert(dataOrNode interface{}, index uint) {
	switch dataOrNode := dataOrNode.(type) {
	case *Node[T]:
		temp := l.First_node
		var temp_index uint = 0
		for temp_index != index {
			if temp.HasNext() {
				temp = temp.Next
				temp_index++
			} else {
				return
			}
		}
		dataOrNode.Prev = temp.Prev // [prev] <-  [node]     [temp]
		dataOrNode.Next = temp      // [prev] <-  [node]  -> [temp]
		temp.Prev.Next = dataOrNode // [prev] <-> [node]  -> [temp]
		temp.Prev = dataOrNode      // [prev] <-> [node] <-> [temp]
		l.Len++
	case T:
		node := NewNode(dataOrNode)
		temp := l.First_node
		var temp_index uint = 0
		for temp_index != index {
			if temp.HasNext() {
				temp = temp.Next
				temp_index++
			} else {
				return
			}
		}
		node.Prev = temp.Prev // [prev] <-  [node]     [temp]
		node.Next = temp      // [prev] <-  [node]  -> [temp]
		temp.Prev.Next = node // [prev] <-> [node]  -> [temp]
		temp.Prev = node      // [prev] <-> [node] <-> [temp]
		l.Len++
	}
}

func (l *DoublyLinkedList[T]) Remove(index uint) {
	temp := l.First_node
	var temp_index uint = 0
	for temp_index != index {
		if temp.HasNext() {
			temp = temp.Next
			temp_index++
		} else {
			return
		}
	}
	if temp == l.First_node {
		l.First_node = temp.Next
	}
	if temp == l.Current {
		l.Current = l.First_node
	}
	if temp.HasPrev() && temp.HasNext() {
		temp.Prev.Next = temp.Next // [prev]  -> [next]
		temp.Next.Prev = temp.Prev // [prev] <-> [next]
	} else if temp.HasPrev() {
		temp.Prev.Next = nil
	} else if temp.HasNext() {
		temp.Next.Prev = nil
	}

	l.Len--
}

func (l *DoublyLinkedList[T]) RemoveLeft() {
	temp := l.First_node
	if temp != nil {
		if temp.HasNext() {
			l.First_node = temp.Next
			if l.Current == temp {
				l.Current = l.First_node
			}
			temp.Next.Prev = nil
		} else {
			l.First_node = nil
			l.Current = nil
		}
		l.Len--
	}
}

func (l *DoublyLinkedList[T]) RemoveRight() {
	temp := l.Last()
	if temp != nil {
		if temp.HasPrev() {
			temp.Prev.Next = nil
			if l.Current == temp {
				l.Current = temp.Prev
			}
		}
		if l.First_node == temp {
			l.First_node = nil
			l.Current = l.First_node
		}
		l.Len--
	}
}

// Check that there is a such value in the dll
func (l *DoublyLinkedList[T]) Contains(data T) bool {
	if l.Len == 0 {
		return false
	}
	temp := l.First_node
	if temp.Data == data {
		return true
	}
	for temp.HasNext() {
		temp = temp.Next
		if temp.Data == data {
			return true
		}
	}
	return false
}

// Get node by the given index
func (l *DoublyLinkedList[T]) Get(index uint) (node *Node[T], err error) {
	if l.Len == 0 {
		return nil, errors.New("No element in the list")
	}
	temp := l.First_node
	var temp_index uint = 0
	for temp_index != index {
		if temp.HasNext() {
			temp = temp.Next
			temp_index++
		} else {
			return nil, errors.New("Index out of range")
		}

	}
	return temp, nil
}

// Get the index of the given data
func (l *DoublyLinkedList[T]) GetIndex(data T) (index uint, err error) {
	if !l.Contains(data) {
		return 0, errors.New("No such element in the list")
	}
	temp := l.First_node
	var temp_index uint = 0
	if temp.Data == data {
		return temp_index, nil
	}
	for temp.HasNext() {
		temp = temp.Next
		temp_index++
		if temp.Data == data {
			return temp_index, nil
		}
	}
	return 0, errors.New("No such data in the list")
}

// Get new list where all Node link are reversed. NODES WILL NOT BE COPIED and only their links will be reversed.
//
// Old list will be unusable after reversing. Must be used as dll = dll.GetReversed()
func (l *DoublyLinkedList[T]) GetReversed() *DoublyLinkedList[T] {
	new_list := NewDoublyLinkedList[T]()
	temp := l.First_node
	new_list.AppendLeft(temp)
	for temp.HasNext() {
		temp = temp.Next
		new_list.AppendLeft(temp)
	}
	new_list.Len = l.Len
	new_list.Current = l.Current
	return new_list
}

// Stringer method.
// Outputs formatted string representation of the list.
func (l *DoublyLinkedList[T]) String() string {	
	var str string = "DoublyLinkedList:"
	if l.First_node != nil {
		str = str + "\nFirst_node:\n" + l.First_node.String()
	}
	if l.Current != nil {
		str = str + "\nCurrent:\n" + l.Current.String()
	}
	str = str + "\nLen:" + string(l.Len)
	str = str + "\nNodes:"
	temp := l.First_node
	if temp == l.Current {
		str = str + "\n" + "^<0>" +temp.String() + "<0>^"
	} else {
		str = str + "\n" + "<0>" +temp.String()	+ "<0>"
	}
	str = str + "\n" + temp.String()
	for temp.HasNext() {
		temp = temp.Next
		if temp == l.Current {
			str = str + "\n" + "^" + temp.String() + "^"
		}
	}
	return str
}
