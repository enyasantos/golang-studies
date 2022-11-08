package main

import (
	"fmt"
	"linkedlist/linkedlist"
)

func main() {
	list := linkedlist.NewLinkedList()
	list.Append("Item 1")
	list.Append("Item 2")
	list.Append("Item 3")
	list.Append("Item 4")
	fmt.Println(list)

	search := list.Search("Item 3")
	fmt.Println(search.Value)

	list.Delete("Item 2")
	fmt.Println(list)

}
