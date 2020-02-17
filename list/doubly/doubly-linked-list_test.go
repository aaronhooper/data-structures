package doubly

import (
	"testing"
)

func TestCreate(t *testing.T) {
	var expectation bool
	result := Create()

	expectation = result.head == nil && result.tail == nil
	if !expectation {
		t.Error("Expected List.head and List.tail to be nil")
	}

	expectation = result.len == 0
	if !expectation {
		t.Error("Expected List.len to equal 0")
	}
}

func TestInsertFirst(t *testing.T) {
	var expectation bool
	list := Create()
	list.InsertFirst(10)

	expectation = list.head.data == 10
	if !expectation {
		t.Error("Expected List.head.data to equal 10")
	}
	expectation = list.tail.data == 10
	if !expectation {
		t.Error("Expected List.tail.data to equal 10")
	}
	expectation = list.len == 1
	if !expectation {
		t.Error("Expected List.len to equal 1")
	}

	list.InsertFirst(45)

	expectation = list.head.data == 45
	if !expectation {
		t.Error("Expected List.head.data to equal 45")
	}
	expectation = list.tail.data == 10
	if !expectation {
		t.Error("Expected List.tail.data to equal 10")
	}
	expectation = list.len == 2
	if !expectation {
		t.Error("Expected List.len to equal 2")
	}
}

func TestInsertLast(t *testing.T) {
	var expectation bool
	list := Create()
	list.InsertLast(245)

	expectation = list.head.data == 245 && list.tail.data == 245
	if !expectation {
		t.Error("Expected head.data and tail.data to equal 245")
	}
	expectation = list.len == 1
	if !expectation {
		t.Error("Expected len to equal 1")
	}

	list.InsertLast(299)

	expectation = list.head.data == 245
	if !expectation {
		t.Error("Expected head.data to equal 245")
	}
	expectation = list.tail.data == 299
	if !expectation {
		t.Error("Expected tail.data to equal 299")
	}
	expectation = list.len == 2
	if !expectation {
		t.Error("Expected len to equal 2")
	}
}

func TestInsertAt(t *testing.T) {
	var expectation bool
	list := Create()
	result := list.InsertAt(3, 284)

	expectation = result != nil
	if !expectation {
		t.Error("Expected InsertAt to return error")
	}

	result = list.InsertAt(0, 76)

	expectation = result == nil
	if !expectation {
		t.Error("Expected InsertAt to return nil")
	}

	result = list.InsertAt(1, 923)

	expectation = result == nil
	if !expectation {
		t.Error("Expected InsertAt to return nil")
	}

	result = list.InsertAt(1, 55)

	expectation = list.len == 3
	if !expectation {
		t.Error("Expected len to equal 3")
	}
	expectation = list.head.next.data == 55
	if !expectation {
		t.Error("Expected head.next.data to equal 55")
	}
	expectation = list.tail.prev.data == 55
	if !expectation {
		t.Error("Expected tail.prev.data to equal 55")
	}
}

func TestSearch(t *testing.T) {
	var expectation bool
	list := Create()

	pos, err := list.Search(43)

	expectation = err != nil
	if !expectation {
		t.Error("Expected Search to return an error")
	}

	list.InsertLast(245)
	list.InsertLast(43)
	list.InsertLast(2966)
	list.InsertLast(1)

	pos, err = list.Search(245)

	expectation = pos == 0 && err == nil
	if !expectation {
		t.Error("Expected Search to return 0 and no error")
	}

	pos, err = list.Search(1)

	expectation = pos == 3 && err == nil
	if !expectation {
		t.Error("Expected Search to return 3 and no error")
	}

	pos, err = list.Search(2966)

	expectation = pos == 2 && err == nil
	if !expectation {
		t.Error("Expected Search to return 3 and no error")
	}

	pos, err = list.Search(55)

	expectation = err != nil
	if !expectation {
		t.Error("Expected Search to return an error")
	}
}

func TestDataAt(t *testing.T) {
	var expectation bool
	list := Create()

	data, err := list.DataAt(4)

	expectation = err != nil
	if !expectation {
		t.Error("Expected DataAt to return an error")
	}

	list.InsertLast(256)
	list.InsertLast(91)
	list.InsertLast(52)
	data, err = list.DataAt(0)

	expectation = data == 256 && err == nil
	if !expectation {
		t.Error("Expected DataAt to return 256 and no error")
	}

	data, err = list.DataAt(2)

	expectation = data == 52 && err == nil
	if !expectation {
		t.Error("Expected DataAt to return 52 and no error")
	}

	data, err = list.DataAt(1)

	expectation = data == 91 && err == nil
	if !expectation {
		t.Error("Expected DataAt to return 91 and no error")
	}

	list.InsertLast(1010)
	list.InsertLast(1925)
	data, err = list.DataAt(1)

	expectation = data == 91 && err == nil
	if !expectation {
		t.Error("Expected DataAt to return 91 and no error")
	}
}

func TestRemoveFirst(t *testing.T) {
	var expectation bool
	list := Create()

	data, err := list.RemoveFirst()

	expectation = err != nil
	if !expectation {
		t.Error("Expected RemoveFirst to return an error")
	}

	list.InsertLast(42)
	list.InsertLast(21)

	data, err = list.RemoveFirst()

	expectation = data == 42 && err == nil
	if !expectation {
		t.Error("Expected RemoveFirst to return 42 and no error")
	}
	expectation = list.head.data == 21 && list.tail.data == 21
	if !expectation {
		t.Error("Expected head.data and tail.data to equal 21")
	}

	data, err = list.RemoveFirst()

	expectation = data == 21 && err == nil
	if !expectation {
		t.Error("Expected RemoveFirst to return 21 and no error")
	}
	expectation = list.head == nil && list.tail == nil
	if !expectation {
		t.Error("Expected List head and tail to be nil")
	}
}

func TestRemoveLast(t *testing.T) {
	var expectation bool
	list := Create()

	data, err := list.RemoveLast()

	expectation = err != nil
	if !expectation {
		t.Error("Expected RemoveLast to return an error")
	}

	list.InsertFirst(42)
	list.InsertFirst(21)

	data, err = list.RemoveLast()

	expectation = data == 42 && err == nil
	if !expectation {
		t.Error("Expected RemoveLast to return 42 and no error")
	}
	expectation = list.head.data == 21 && list.tail.data == 21
	if !expectation {
		t.Error("Expected head.data and tail.data to equal 21")
	}

	data, err = list.RemoveLast()

	expectation = data == 21 && err == nil
	if !expectation {
		t.Error("Expected RemoveLast to return 21 and no error")
	}
	expectation = list.head == nil && list.tail == nil
	if !expectation {
		t.Error("Expected List head and tail to be nil")
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

	list.InsertLast(35)
	data, err = list.RemoveAt(0)

	expectation = data == 35 && err == nil
	if !expectation {
		t.Error("Expected RemoveAt to return 35 and no error")
	}

	list.InsertLast(47)
	list.InsertLast(56)
	data, err = list.RemoveAt(1)

	expectation = data == 56 && err == nil
	if !expectation {
		t.Error("Expected RemoveAt to return 56 and no error")
	}

	list.InsertLast(10)
	list.InsertLast(99)
	data, err = list.RemoveAt(1)

	expectation = data == 10 && err == nil
	if !expectation {
		t.Error("Expected RemoveAt to return 10 and no error")
	}

	list.InsertLast(256)
	list.InsertLast(300)
	list.InsertLast(2)
	data, err = list.RemoveAt(1)

	expectation = data == 99 && err == nil
	if !expectation {
		t.Error("Expected RemoveAt to return 99 and no error")
	}
}
