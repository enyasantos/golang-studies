package main

import (
	"data-structures/stack"
	"fmt"
)

func main() {
	stack := stack.NewStack(5)

	stack.Push("Item 1")
	stack.Push("Item 2")
	stack.Push("Item 3")
	stack.Push("Item 4")
	fmt.Println(stack)
}
