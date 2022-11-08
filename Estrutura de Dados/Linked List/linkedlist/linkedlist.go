package linkedlist

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Append(value string) {
	node := Node{Value: value}

	if l.Head == nil {
		l.Head = &node
	}

	if l.Tail != nil {
		l.Tail.Next = &node
		node.Prev = l.Tail.Next
	}

	l.Tail = &node
}

func (l *LinkedList) Search(value string) *Node {
	node := l.Head
	for node != nil {
		if node.Value == value {
			return node
		}
		node = node.Next
	}
	return nil
}

func (l *LinkedList) Delete(value string) {
	if l.Tail.Value == value {
		l.Tail = l.Head.Prev
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	prev := l.Head
	node := l.Head.Next
	for node != nil {
		if node.Value == value {
			prev.Next = node.Next
			node.Next = prev
			break
		}
		prev = node
		node = node.Next
	}
}

func (l *LinkedList) String() string {
	str := ""
	node := l.Head
	for node != nil {
		str = str + " -> " + node.Value
		node = node.Next
	}
	return str
}

type Node struct {
	Value string
	Next  *Node
	Prev  *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
