package singly

import (
	"fmt"
)

type NodeData interface{}

type node struct {
	data NodeData
	next *node
}

type List struct {
	head *node
	len  uint
}

// Create returns an empty singly linked list.
func Create() List {
	return List{nil, 0}
}

func (list *List) Len() uint {
	return list.len
}

// InsertFirst inserts `data` at the beginning of the list.
func (list *List) InsertFirst(data NodeData) error {
	oldHead := list.head
	newHead := node{data, oldHead}
	list.head = &newHead
	list.len++
	return nil
}

// InsertAt inserts `data` at the position `pos` in the list.
// Returns an error if `pos` is an invalid position.
func (list *List) InsertAt(pos uint, data NodeData) error {
	if pos > list.len {
		return fmt.Errorf("Cannot insert at out of bounds position %v", pos)
	}

	if pos == 0 {
		return list.InsertFirst(data)
	}

	var trav *node

	for i := uint(0); i < pos; i++ {
		if i == uint(0) {
			trav = list.head
		} else {
			trav = trav.next
		}
	}

	newNode := node{data, trav.next}
	trav.next = &newNode
	list.len++
	return nil
}

// DataAt returns the data stored at `pos`, or an error if `pos` is an invalid
// position or the list is empty.
func (list *List) DataAt(pos uint) (NodeData, error) {
	if list.len == 0 {
		return 0, fmt.Errorf("Cannot read from empty list")
	}

	if pos >= list.len {
		return 0, fmt.Errorf("Cannot read from out of bounds position %v", pos)
	}

	var trav *node
	for i := uint(0); i <= pos; i++ {
		if i == uint(0) {
			trav = list.head
		} else {
			trav = trav.next
		}
	}

	return trav.data, nil
}

// Search returns the position of the first occurrence of `data`, or an error
// if none was found or the list is empty.
func (list *List) Search(data NodeData) (uint, error) {
	if list.len == 0 {
		return 0, fmt.Errorf("Cannot search empty list")
	}

	head := list.head
	trav := head

	for i := uint(0); i < list.len; i++ {
		if data == trav.data {
			return i, nil
		}
		trav = trav.next
	}

	return 0, fmt.Errorf("%v not found in list", data)
}

// RemoveFirst removes the first element in the list and returns it, or
// an error if if the list is empty.
func (list *List) RemoveFirst() (NodeData, error) {
	if list.len == 0 {
		return 0, fmt.Errorf("Cannot remove from empty list")
	}

	data := list.head.data
	next := list.head.next
	list.head = next
	list.len--
	return data, nil
}

// RemoveAt removes the int at position `pos` in the list and returns it.
// Returns an error if `pos` is out of bounds.
func (list *List) RemoveAt(pos uint) (NodeData, error) {
	if pos >= list.len {
		return 0, fmt.Errorf("Cannot remove out of bounds position in list")
	}

	if pos == 0 {
		return list.RemoveFirst()
	}

	var prev *node
	trav := list.head

	for i := uint(0); i < pos; i++ {
		prev = trav
		trav = trav.next
	}

	next := trav.next
	prev.next = next
	list.len--
	return trav.data, nil
}

func (list List) String() string {
	if list.len == 0 {
		return "[]"
	}

	output := "["
	var trav *node

	for i := uint(0); i < list.len; i++ {
		if i == uint(0) {
			trav = list.head
		} else {
			trav = trav.next
		}

		output = fmt.Sprintf("%s%v", output, trav.data)

		if i == list.len-1 {
			output += "]"
		} else {
			output += ", "
		}
	}

	return output
}
