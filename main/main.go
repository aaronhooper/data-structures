package main

import (
	"fmt"

	"github.com/aaronhooper/data-structures/list/doubly"
	"github.com/aaronhooper/data-structures/list/singly"
)

func main() {
	fmt.Println("--- SINGLY LINKED LIST ---")
	var singlyList = singly.Create(23)
	fmt.Println(singlyList)
	singlyList.Insert(46, 0)
	fmt.Println(singlyList)
	singlyList.Insert(193, 2)
	fmt.Println(singlyList)

	singlyList.Append(99)
	singlyList.Append(2)
	singlyList.Append(714)
	fmt.Println(singlyList)

	fmt.Println(singlyList.Search(46))
	fmt.Println(singlyList.Search(23))
	fmt.Println(singlyList.Search(193))
	fmt.Println(singlyList.Search(2334))

	fmt.Println(singlyList.DataAt(0))
	fmt.Println(singlyList.DataAt(1))
	fmt.Println(singlyList.DataAt(2))
	fmt.Println(singlyList.DataAt(3))

	singlyList.Remove(0)
	fmt.Println(singlyList)
	singlyList.Remove(1)
	fmt.Println(singlyList)

	fmt.Println("--- DOUBLY LINKED LIST ---")
	doublyList := doubly.Create(2)
	fmt.Println(doublyList)
	doublyList.Insert(10, 0)
	fmt.Println(doublyList)
	doublyList.Insert(245, 2)
	fmt.Println(doublyList)
	// => [10, 2, 245]
}
