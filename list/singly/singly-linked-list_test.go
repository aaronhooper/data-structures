package singly

import (
	"testing"
)

func TestCreate(t *testing.T) {
	var expectation bool
	result := Create()

	expectation = result.head == nil
	if !expectation {
		t.Error("Expected List.head to be nil")
	}

	expectation = result.len == 0
	if !expectation {
		t.Error("Expected List.len to equal 0")
	}
}

func TestInsertFirst(t *testing.T) {
	var expectation bool
	list := Create()

	list.InsertFirst(32)

	expectation = list.head.data == 32 && list.len == 1
	if !expectation {
		t.Error("Expected head.data to equal 32 and len to equal 1")
	}
}

func TestInsertAt(t *testing.T) {
	var expectation bool
	list := Create()

	err := list.InsertAt(1, 34)

	expectation = err != nil
	if !expectation {
		t.Error("Expected InsertAt to return an error")
	}

	list.InsertAt(0, 38)

	expectation = list.head.data == 38
	if !expectation {
		t.Error("Expected head.data to equal 38")
	}

	list.InsertAt(1, 99)

	expectation = list.head.next.data == 99
	if !expectation {
		t.Error("Expected head.next.data to equal 99")
	}

	list.InsertAt(2, 456)

	expectation = list.head.next.next.data == 456
	if !expectation {
		t.Error("Expected head.next.next.data to equal 456")
	}

}

func TestDataAt(t *testing.T) {
	var expectation bool
	list := Create()

	data, err := list.DataAt(0)

	expectation = err != nil
	if !expectation {
		t.Error("Expected DataAt to return an error")
	}

	list.InsertFirst(31)
	data, err = list.DataAt(1)

	expectation = err != nil
	if !expectation {
		t.Error("Expected DataAt to return an error")
	}

	data, err = list.DataAt(0)

	expectation = data == 31
	if !expectation {
		t.Error("Expected DataAt to return 31")
	}

	list.InsertFirst(31)
	data, err = list.DataAt(1)

	expectation = data == 31
	if !expectation {
		t.Error("Expected DataAt to return 31")
	}
}

func TestSearch(t *testing.T) {
	var expectation bool
	list := Create()

	pos, err := list.Search(76)

	expectation = err != nil
	if !expectation {
		t.Error("Expected Search to return an error")
	}

	list.InsertFirst(34)
	pos, err = list.Search(34)

	expectation = pos == 0
	if !expectation {
		t.Error("Expected Search to return 0")
	}

	list.InsertAt(1, 11)
	pos, err = list.Search(11)

	expectation = pos == 1
	if !expectation {
		t.Error("Expected Search to return 1")
	}

	pos, err = list.Search(289)

	expectation = err != nil
	if !expectation {
		t.Error("Expected Search to return an error")
	}
}

func TestRemoveFirst(t *testing.T) {
	var expectation bool
	list := Create()
	data, err := list.RemoveFirst()

	expectation = err != nil
	if !expectation {
		t.Error("Expected RemoveFirst to return error")
	}

	list.InsertFirst(25)
	data, err = list.RemoveFirst()

	expectation = data == 25
	if !expectation {
		t.Error("Expected RemoveFirst to return 25")
	}
}

func TestRemoveAt(t *testing.T) {
	var expectation bool
	list := Create()

	data, err := list.RemoveAt(0)

	expectation = err != nil
	if !expectation {
		t.Error("Expected RemoveAt to return an error")
	}

	list.InsertFirst(34)
	data, err = list.RemoveAt(0)

	expectation = data == 34
	if !expectation {
		t.Error("Expected RemoveAt to return 34")
	}

	list.InsertFirst(99)
	list.InsertFirst(168)
	data, err = list.RemoveAt(1)

	expectation = data == 99
	if !expectation {
		t.Error("Expected RemoveAt to return 99")
	}
}
