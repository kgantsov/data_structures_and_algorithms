package linked_list

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

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList(args ...int) *LinkedList {
	list := new(LinkedList)

	return list
}

func (l *LinkedList) append(value int) {
	if l.head == nil {
		l.head = NewNode(value)
		l.tail = l.head
	} else {
		l.tail.next = NewNode(value)
		l.tail = l.tail.next
	}
}

func (l *LinkedList) search(value int) (foundValue int, ok bool) {
	node := l.head

	if node == nil {
		return 0, false
	}

	for node.next != nil {
		if node.value == value {
			return node.value, true
		}
		node = node.next
	}
	return 0, false
}

func (l *LinkedList) index(value int) (index int, ok bool) {
	node := l.head
	index = 0

	if node == nil {
		return index, false
	}

	for node.next != nil {
		if node.value == value {
			return index, true
		}
		node = node.next
		index++
	}
	return index, false
}

func (l *LinkedList) String() string {
	node := l.head

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

func (l *LinkedList) Len() int {
	node := l.head

	var length int

	if node == nil {
		return 0
	}

	for node.next != nil {
		length++

		node = node.next
	}
	length++
	return length
}
