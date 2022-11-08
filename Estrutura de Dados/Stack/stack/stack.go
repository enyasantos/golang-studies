package stack

type Stack struct {
	Tail  *Node
	Size  uint
	Count uint
}

func (s *Stack) Push(value string) {
	if s.Count == s.Size {
		panic("Stack overflow")
	}

	node := Node{Value: value}
	if s.Tail != nil {
		node.Next = s.Tail
	}

	s.Tail = &node
	s.Count += 1
}

func (s *Stack) Pop() string {
	if s.Tail == nil {
		return ""
	}
	s.Count -= 1
	node := s.Tail
	s.Tail = node.Next

	return node.Value
}

func (s *Stack) String() string {
	str := ""
	node := s.Tail
	for node != nil {
		str = str + " -> " + node.Value
		node = node.Next
	}

	return str
}

type Node struct {
	Value string
	Next  *Node
}

func NewStack(size uint) *Stack {
	return &Stack{Size: size}
}
