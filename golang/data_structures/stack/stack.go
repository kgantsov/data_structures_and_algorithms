package stack

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	node := new(Node)
	node.value = value

	return node
}

type Stack struct {
	top *Node
}

func NewStack(args ...int) *Stack {
	list := new(Stack)

	for _, v := range args {
		list.push(v)
	}

	return list
}

func (s *Stack) String() string {
	node := s.top

	var values []string

	if node == nil {
		return "[]"
	}

	for node.next != nil {
		values = append(values, strconv.Itoa(node.value))

		node = node.next
	}
	values = append(values, strconv.Itoa(node.value))

	return fmt.Sprintf("[%s]", strings.Join(values, ", "))
}

func (s *Stack) push(value int) {
	newTop := NewNode(value)

	if s.top == nil {
		s.top = newTop
	} else {
		newTop.next = s.top
		s.top = newTop
	}
}

func (s *Stack) pop() (value int, ok bool) {

	if s.top == nil {
		return 0, false
	} else {
		node := s.top
		s.top = node.next
		return node.value, true
	}
}

func (s *Stack) peek() (value int, ok bool) {
	if s.top == nil {
		return 0, false
	} else {
		return s.top.value, true
	}
}
