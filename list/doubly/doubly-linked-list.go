package doubly

import (
	"fmt"
)

type NodeData interface{}

type node struct {
	data NodeData
	prev *node
	next *node
}

type List struct {
	head *node
	tail *node
	len  uint
}

// Create returns an empty doubly linked list.
func Create() List {
	return List{nil, nil, 0}
}

// InsertFirst inserts `data` into the beginning of the list.
func (list *List) InsertFirst(data NodeData) error {
	oldHead := list.head
	newHead := node{data, nil, oldHead}

	if list.len == 0 {
		list.tail = &newHead
	} else {
		oldHead.prev = &newHead
	}

	list.head = &newHead
	list.len++
	return nil
}

// InsertLast inserts `data` at the end of the list.
func (list *List) InsertLast(data NodeData) error {
	if list.len == 0 {
		return list.InsertFirst(data)
	}

	oldTail := list.tail
	newTail := node{data, oldTail, nil}
	newTail.prev.next = &newTail
	list.tail = &newTail
	list.len++
	return nil
}

// InsertAt inserts `data` at position `pos` in the list. Returns an error
// if `pos` is out of bounds.
func (list *List) InsertAt(pos uint, data NodeData) error {
	if pos > list.len {
		return fmt.Errorf("Cannot insert at out of bounds position %v", pos)
	}

	if pos == 0 {
		return list.InsertFirst(data)
	}

	if pos == list.len {
		return list.InsertLast(data)
	}

	trav := list.head

	for i := uint(0); i < pos; i++ {
		trav = trav.next
	}

	newNode := node{data, trav.prev, trav}
	trav.prev.next = &newNode
	trav.prev = &newNode
	list.len++
	return nil
}

// DataAt returns the data at the position `pos` or an error if `pos`
// is out of bounds.
func (list *List) DataAt(pos uint) (NodeData, error) {
	if pos >= list.len {
		return 0, fmt.Errorf(
			"Cannot read from out of bounds position %v", pos)
	}

	if pos == 0 {
		return list.head.data, nil
	}

	if pos == list.len-1 {
		return list.tail.data, nil
	}

	var trav *node
	posInFirstHalf := pos < list.len/2

	if posInFirstHalf {
		trav = list.head

		for i := uint(0); i < pos; i++ {
			trav = trav.next
		}
	} else {
		trav = list.tail

		for i := list.len - 1; i > pos; i-- {
			trav = trav.prev
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

// RemoveFirst removes the first element in the list and returns it.
// Returns an error if the list is empty.
func (list *List) RemoveFirst() (NodeData, error) {
	if list.len == 0 {
		return 0, fmt.Errorf("Cannot remove from empty list")
	}

	headData := list.head.data
	list.head = list.head.next

	if list.len == 1 {
		list.tail = nil
	} else {
		list.head.prev = nil
	}

	list.len--
	return headData, nil
}

// RemoveLast removes the last element in the list and returns it.
// Returns an error if the list is empty.
func (list *List) RemoveLast() (NodeData, error) {
	if list.len == 0 {
		return 0, fmt.Errorf("Cannot remove from empty list")
	}

	if list.len == 1 {
		return list.RemoveFirst()
	}

	tailData := list.tail.data
	list.tail = list.tail.prev
	list.tail.next = nil
	list.len--
	return tailData, nil
}

// RemoveAt removes the element located at `pos` in the list and returns it.
// Returns an error if `pos` is out of bounds.
func (list *List) RemoveAt(pos uint) (NodeData, error) {
	if pos >= list.len {
		return 0, fmt.Errorf(
			"Cannot remove from out of bounds position %v", pos)
	}

	if pos == 0 {
		return list.RemoveFirst()
	}

	if pos == list.len-1 {
		return list.RemoveLast()
	}

	var trav *node
	posInFirstHalf := pos < list.len/2

	if posInFirstHalf {
		trav = list.head

		for i := uint(0); i < pos; i++ {
			trav = trav.next
		}
	} else {
		trav = list.tail

		for i := list.len - 1; i > pos; i-- {
			trav = trav.prev
		}
	}

	travData := trav.data
	prev := trav.prev
	next := trav.next
	prev.next = next
	next.prev = prev
	list.len--
	return travData, nil
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

		output += fmt.Sprintf("%v", trav.data)

		if i == list.len-1 {
			output += "]"
		} else {
			output += ", "
		}
	}

	return output
}
