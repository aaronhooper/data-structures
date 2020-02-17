package main

import (
	"fmt"

	"github.com/aaronhooper/data-structures/list/doubly"
	"github.com/aaronhooper/data-structures/list/singly"
)

func main() {
	fmt.Println("--- SINGLY LINKED LIST ---")
	var singlyList = singly.Create()
	fmt.Println(singlyList)
	// => []

	fmt.Println("***")
	singlyList.InsertFirst(23)
	singlyList.InsertFirst(46)
	singlyList.InsertAt(2, 193)
	fmt.Println(singlyList)
	// => [46, 23, 193]

	fmt.Println("***")
	singlyList.InsertFirst(99)
	singlyList.InsertFirst(2)
	singlyList.InsertFirst(714)
	fmt.Println(singlyList)
	// => [714, 2, 99, 46, 23, 193]

	fmt.Println("***")
	fmt.Println(singlyList.Search(46))
	fmt.Println(singlyList.Search(23))
	fmt.Println(singlyList.Search(193))
	fmt.Println(singlyList.Search(2334))
	// => 3, 4, 5, error

	fmt.Println("***")
	fmt.Println(singlyList.DataAt(0))
	fmt.Println(singlyList.DataAt(1))
	fmt.Println(singlyList.DataAt(2))
	fmt.Println(singlyList.DataAt(3))
	// => 714, 2, 99, 46

	fmt.Println("***")
	fmt.Println(singlyList.RemoveFirst())
	fmt.Println(singlyList.RemoveAt(1))
	// => 714, 99

	fmt.Println("***")
	fmt.Println(singlyList.RemoveFirst())
	fmt.Println(singlyList.RemoveFirst())
	fmt.Println(singlyList.RemoveFirst())
	fmt.Println(singlyList.RemoveFirst())
	// => 2, 46, 23, 193

	fmt.Println("***")
	fmt.Println(singlyList)
	// => []

	fmt.Println("***")
	fmt.Println(singlyList.DataAt(0))
	// => error

	fmt.Println("***")
	fmt.Println(singlyList.Search(193))
	// => error

	fmt.Println("***")
	fmt.Println(singlyList.RemoveFirst())
	// => error

	fmt.Println("--- DOUBLY LINKED LIST ---")
	doublyList := doubly.Create()
	fmt.Println(doublyList)
	// => []

	fmt.Println("***")
	doublyList.InsertFirst(10)
	doublyList.InsertLast(245)
	doublyList.InsertAt(1, 2)
	fmt.Println(doublyList)
	// => [10, 2, 245]

	fmt.Println("***")
	doublyList.InsertAt(1, 2966)
	doublyList.InsertAt(0, 43)
	fmt.Println(doublyList)
	// => [43, 10, 2966, 2, 245]

	fmt.Println("***")
	fmt.Println(doublyList.Search(245))
	fmt.Println(doublyList.Search(43))
	fmt.Println(doublyList.Search(2966))
	fmt.Println(doublyList.Search(1))
	// => 4, 0, 2, error

	fmt.Println("***")
	fmt.Println(doublyList.DataAt(0))
	fmt.Println(doublyList.DataAt(1))
	fmt.Println(doublyList.DataAt(2))
	fmt.Println(doublyList.DataAt(3))
	fmt.Println(doublyList.DataAt(4))
	// => 43, 10, 2966, 2, 245

	fmt.Println("***")
	doublyList.RemoveFirst()
	doublyList.RemoveLast()
	doublyList.RemoveAt(1)
	fmt.Println(doublyList)
	// => [10, 2]

	fmt.Println("***")
	doublyList.RemoveLast()
	doublyList.RemoveLast()
	fmt.Println(doublyList)
	// => []

	fmt.Println("***")
	fmt.Println(doublyList.RemoveFirst())
	// => error
}
