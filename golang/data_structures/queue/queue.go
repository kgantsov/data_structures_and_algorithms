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

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue(args ...int) *Queue {
	list := new(Queue)

	for _, v := range args {
		list.enqueue(v)
	}

	return list
}

func (q *Queue) String() string {
	node := q.head

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

func (q *Queue) enqueue(value int) {
	newTail := NewNode(value)

	if q.head == nil {
		q.head = newTail
		q.tail = q.head
	} else {
		q.tail.next = newTail
		q.tail = newTail
	}
}

func (q *Queue) dequeue() (value int, ok bool) {

	if q.head == nil {
		return 0, false
	} else {
		node := q.head
		q.head = node.next
		return node.value, true
	}
}

func (q *Queue) peek() (value int, ok bool) {
	if q.head == nil {
		return 0, false
	} else {
		return q.head.value, true
	}
}
