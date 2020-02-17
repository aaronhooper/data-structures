package singly

import (
	"fmt"
	"strconv"
)

type node struct {
	data int
	next *node
}

type list struct {
	head *node
	len  uint
}

// Create returns an empty singly linked list.
func Create() list {
	return list{nil, 0}
}

// InsertFirst inserts the integer `data` at the beginning of the list.
func (l *list) InsertFirst(data int) error {
	oldHead := l.head
	newHead := node{data, oldHead}
	l.head = &newHead
	l.len++
	return nil
}

// InsertAt inserts the integer `data` at the position `pos` in the list.
// Returns an error if `pos` is an invalid position.
func (l *list) InsertAt(pos uint, data int) error {
	if pos > l.len {
		return fmt.Errorf("Cannot insert at out of bounds position %v", pos)
	}

	if pos == 0 {
		return l.InsertFirst(data)
	}

	var trav *node

	for i := uint(0); i < pos; i++ {
		if i == uint(0) {
			trav = l.head
		} else {
			trav = trav.next
		}
	}

	newNode := node{data, trav.next}
	trav.next = &newNode
	l.len++
	return nil
}

// DataAt returns the data stored at `pos`, or an error if `pos` is an invalid
// position or the list is empty.
func (l *list) DataAt(pos uint) (int, error) {
	if l.len == 0 {
		return 0, fmt.Errorf("Cannot read from empty list")
	}

	if pos >= l.len {
		return 0, fmt.Errorf("Cannot read from out of bounds position %v", pos)
	}

	var trav *node
	for i := uint(0); i <= pos; i++ {
		if i == uint(0) {
			trav = l.head
		} else {
			trav = trav.next
		}
	}

	return trav.data, nil
}

// Search returns the position of the first occurrence of `data`, or an error
// if none was found or the list is empty.
func (l *list) Search(data int) (uint, error) {
	if l.len == 0 {
		return 0, fmt.Errorf("Cannot search empty list")
	}

	head := l.head
	trav := head

	for i := uint(0); i < l.len; i++ {
		if data == trav.data {
			return i, nil
		}
		trav = trav.next
	}

	return 0, fmt.Errorf("%v not found in list", data)
}

// RemoveFirst removes the first element in the list and returns it, or
// an error if if the list is empty.
func (l *list) RemoveFirst() (int, error) {
	if l.len == 0 {
		return 0, fmt.Errorf("Cannot remove from empty list")
	}

	data := l.head.data
	next := l.head.next
	l.head = next
	l.len--
	return data, nil
}

// RemoveAt removes the int at position `pos` in the list and returns it.
// Returns an error if `pos` is out of bounds.
func (l *list) RemoveAt(pos uint) (int, error) {
	if pos >= l.len {
		return 0, fmt.Errorf("Cannot remove out of bounds position in list")
	}

	if pos == 0 {
		return l.RemoveFirst()
	}

	var prev *node
	trav := l.head

	for i := uint(0); i < pos; i++ {
		prev = trav
		trav = trav.next
	}

	next := trav.next
	prev.next = next
	l.len--
	return trav.data, nil
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

		output += strconv.Itoa(trav.data)

		if i == l.len-1 {
			output += "]"
		} else {
			output += ", "
		}
	}

	return output
}
