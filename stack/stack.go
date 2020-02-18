package stack

import (
	"github.com/aaronhooper/data-structures/list/singly"
)

type NodeData interface{}

type Stack struct {
	*singly.List
}

func Create() Stack {
	list := singly.Create()
	return Stack{&list}
}

func (s *Stack) Push(data NodeData) error {
	return s.InsertFirst(data)
}

func (s *Stack) Peek() (NodeData, error) {
	return s.DataAt(0)
}

func (s *Stack) Pop() (NodeData, error) {
	return s.RemoveFirst()
}
