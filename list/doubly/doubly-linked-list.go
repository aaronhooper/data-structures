package doubly

import (
	"fmt"
)

type nodeData interface{}

type node struct {
	data nodeData
	prev *node
	next *node
}

type list struct {
	head *node
	tail *node
	len  uint
}

// Create returns an empty doubly linked list.
func Create() list {
	return list{nil, nil, 0}
}

// InsertFirst inserts `data` into the beginning of the list.
func (l *list) InsertFirst(data nodeData) error {
	oldHead := l.head
	newHead := node{data, nil, oldHead}

	if l.len == 0 {
		l.tail = &newHead
	} else {
		oldHead.prev = &newHead
	}

	l.head = &newHead
	l.len++
	return nil
}

// InsertLast inserts `data` at the end of the list.
func (l *list) InsertLast(data nodeData) error {
	if l.len == 0 {
		return l.InsertFirst(data)
	}

	oldTail := l.tail
	newTail := node{data, oldTail, nil}
	newTail.prev.next = &newTail
	l.tail = &newTail
	l.len++
	return nil
}

// InsertAt inserts `data` at position `pos` in the list. Returns an error
// if `pos` is out of bounds.
func (l *list) InsertAt(pos uint, data nodeData) error {
	if pos > l.len {
		return fmt.Errorf("Cannot insert at out of bounds position %v", pos)
	}

	if pos == 0 {
		return l.InsertFirst(data)
	}

	if pos == l.len {
		return l.InsertLast(data)
	}

	trav := l.head

	for i := uint(0); i < pos; i++ {
		trav = trav.next
	}

	newNode := node{data, trav.prev, trav}
	trav.prev.next = &newNode
	trav.prev = &newNode
	l.len++
	return nil
}

// RemoveFirst removes the first element in the list and returns it.
// Returns an error if the list is empty.
func (l *list) RemoveFirst() (nodeData, error) {
	if l.len == 0 {
		return 0, fmt.Errorf("Cannot remove from empty list")
	}

	headData := l.head.data
	l.head = l.head.next

	if l.len == 1 {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	l.len--
	return headData, nil
}

// RemoveLast removes the last element in the list and returns it.
// Returns an error if the list is empty.
func (l *list) RemoveLast() (nodeData, error) {
	if l.len == 0 {
		return 0, fmt.Errorf("Cannot remove from empty list")
	}

	if l.len == 1 {
		return l.RemoveFirst()
	}

	tailData := l.tail.data
	l.tail = l.tail.prev
	l.tail.next = nil
	l.len--
	return tailData, nil
}

// RemoveAt removes the element located at `pos` in the list and returns it.
// Returns an error if `pos` is out of bounds.
func (l *list) RemoveAt(pos uint) (nodeData, error) {
	if pos >= l.len {
		return 0, fmt.Errorf(
			"Cannot remove from out of bounds position %v", pos)
	}

	if pos == 0 {
		return l.RemoveFirst()
	}

	if pos == l.len-1 {
		return l.RemoveLast()
	}

	var trav *node
	posInFirstHalf := pos < l.len/2

	if posInFirstHalf {
		trav = l.head

		for i := uint(0); i < pos; i++ {
			trav = trav.next
		}
	} else {
		trav = l.tail

		for i := l.len - 1; i > pos; i-- {
			trav = trav.prev
		}
	}

	travData := trav.data
	prev := trav.prev
	next := trav.next
	prev.next = next
	next.prev = prev
	l.len--
	return travData, nil
}

func (l list) String() string {
	if l.len == 0 {
		return "[]"
	}

	output := "["
	var trav *node

	for i := uint(0); i < l.len; i++ {
		if i == uint(0) {
			trav = l.head
		} else {
			trav = trav.next
		}

		output += fmt.Sprintf("%v", trav.data)

		if i == l.len-1 {
			output += "]"
		} else {
			output += ", "
		}
	}

	return output
}
