package doubly

import (
	"fmt"
	"strconv"
)

type node struct {
	data int
	prev *node
	next *node
}

type list struct {
	head node
	tail *node
	len  uint
}

// Create returns a doubly linked list initialized with the value `data`.
func Create(data int) list {
	n := node{data, nil, nil}
	return list{n, &n, 1}
}

func (l *list) Insert(data int, pos uint) error {
	if pos > l.len {
		return fmt.Errorf("Cannot insert at out of bounds position %v", pos)
	}

	if pos == 0 {
		// head must change
		oldHead := l.head
		newHead := node{data, nil, &oldHead}
		oldHead.prev = &newHead
		l.head = newHead
		l.len++
		return nil
	}

	if pos == l.len {
		// tail must change
		newNode := node{data, l.tail, nil}
		newNode.prev.next = &newNode
		l.tail = &newNode
		l.len++
		return nil
	}

	trav := &l.head

	for i := uint(0); i < pos; i++ {
		trav = trav.next
	}

	newNode := node{data, trav.prev, trav}
	trav.prev.next = &newNode
	trav.prev = &newNode
	l.len++
	return nil
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
