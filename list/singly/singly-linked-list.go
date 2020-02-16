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
	head node
	len  uint
}

// Create returns a singly linked list initialized with the value `data`.
func Create(data int) list {
	return list{node{data, nil}, 1}
}

// Insert inserts the integer `data` at the position `pos` in the list.
// Returns an error if `pos` is an invalid position.
func (l *list) Insert(data int, pos uint) error {
	if pos > l.len {
		return fmt.Errorf("Cannot insert at out of bounds position %v", pos)
	}

	if pos == 0 {
		// head must change
		oldHead := l.head
		newHead := node{data, &oldHead}
		l.head = newHead
		l.len++
		return nil
	}

	var trav *node

	for i := uint(0); i < pos; i++ {
		if i == uint(0) {
			trav = &l.head
		} else {
			trav = trav.next
		}
	}

	newNode := node{data, trav.next}
	trav.next = &newNode
	l.len++
	return nil
}

// Append inserts the integer `data` at the head of the list.
func (l *list) Append(data int) error {
	return l.Insert(data, 0)
}

// DataAt returns the data stored at `pos`, or an error if `pos` is an invalid
// position.
func (l *list) DataAt(pos uint) (int, error) {
	if pos >= l.len {
		return 0, fmt.Errorf("Cannot read from out of bounds position %v", pos)
	}

	var trav node
	for i := uint(0); i <= pos; i++ {
		if i == uint(0) {
			trav = l.head
		} else {
			trav = *trav.next
		}
	}

	return trav.data, nil
}

// Search returns the position of the first occurrence of `data`, or an error
// if none was found.
func (l *list) Search(data int) (uint, error) {
	head := l.head
	trav := &head

	for i := uint(0); i < l.len; i++ {
		if data == trav.data {
			return i, nil
		}
		trav = trav.next
	}

	return 0, fmt.Errorf("%v not found in list", data)
}

// Remove removes the int at position `pos` in the list and returns it.
func (l *list) Remove(pos uint) (int, error) {
	trav := &l.head

	if pos == 0 {
		if l.len == 1 {
			return 0, fmt.Errorf("Cannot remove only element in list")
		}

		// head must change
		next := l.head.next
		l.head = *next
		l.len--
		return trav.data, nil
	}

	var prev *node
	for i := uint(0); i < pos; i++ {
		prev = trav
		trav = trav.next

		if trav == nil {
			return 0, fmt.Errorf("Cannot remove out of bounds position in list")
		}
	}

	next := trav.next
	prev.next = next
	l.len--
	return trav.data, nil
}

func (l list) String() string {
	output := "["
	var trav node

	for i := uint(0); i < l.len; i++ {
		if i == uint(0) {
			trav = l.head
		} else {
			trav = *trav.next
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
