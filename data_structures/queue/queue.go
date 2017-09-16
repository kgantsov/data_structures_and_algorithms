package queue

import (
	"fmt"
	"strings"
)

type Node struct {
	value interface{}
	next  *Node
}

func NewNode(value interface{}) *Node {
	node := new(Node)
	node.value = value

	return node
}

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue(args ...interface{}) *Queue {
	list := new(Queue)

	for _, v := range args {
		list.Enqueue(v)
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
		values = append(values, fmt.Sprintf("%v", node.value))

		node = node.next
	}
	values = append(values, fmt.Sprintf("%v", node.value))

	return fmt.Sprintf("[%s]", strings.Join(values, ", "))
}

func (q *Queue) Enqueue(value interface{}) {
	newTail := NewNode(value)

	if q.head == nil {
		q.head = newTail
		q.tail = q.head
	} else {
		q.tail.next = newTail
		q.tail = newTail
	}
}

func (q *Queue) Dequeue() (value interface{}, ok bool) {

	if q.head == nil {
		return 0, false
	} else {
		node := q.head
		q.head = node.next
		return node.value, true
	}
}

func (q *Queue) Peek() (value interface{}, ok bool) {
	if q.head == nil {
		return 0, false
	} else {
		return q.head.value, true
	}
}

func (q *Queue) Empty() bool {
	if q.head == nil {
		return true
	} else {
		return false
	}
}
