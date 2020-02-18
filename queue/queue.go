package queue

import (
	"github.com/aaronhooper/data-structures/list/doubly"
)

type NodeData interface{}

type Queue struct {
	*doubly.List
}

func Create() Queue {
	list := doubly.Create()
	return Queue{&list}
}

func (q *Queue) Enqueue(data NodeData) error {
	return q.InsertFirst(data)
}

func (q *Queue) Dequeue() (NodeData, error) {
	return q.RemoveLast()
}
