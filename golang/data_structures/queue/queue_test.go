package stack

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	myQueue := NewQueue()

	if myQueue.String() != "[]" {
		t.Error("Expected", "[]", ", got ", myQueue.String())
	}

	myQueue = NewQueue(0, 5, 10, 15, 20, 25, 777, 32)

	if myQueue.String() != "[0, 5, 10, 15, 20, 25, 777, 32]" {
		t.Error("Expected", "[0, 5, 10, 15, 20, 25, 777, 32]", ", got ", myQueue.String())
	}

	myQueue = NewQueue(0, 17, -45)

	if myQueue.String() != "[0, 17, -45]" {
		t.Error("Expected", "[0, 17, -45]", ", got ", myQueue.String())
	}
}

func TestQueuePush(t *testing.T) {
	myQueue := NewQueue()
	myQueue.enqueue(10)
	myQueue.enqueue(15)
	myQueue.enqueue(90)
	myQueue.enqueue(35)

	if myQueue.String() != "[10, 15, 90, 35]" {
		t.Error("Expected", "[10, 15, 90, 35]", ", got ", myQueue.String())
	}

	myQueue = NewQueue()
	myQueue.enqueue(77)

	if myQueue.String() != "[77]" {
		t.Error("Expected", "[77]", ", got ", myQueue.String())
	}
	myQueue.enqueue(22)
	myQueue.enqueue(77)
	myQueue.enqueue(123)
	myQueue.enqueue(321)

	if myQueue.String() != "[77, 22, 77, 123, 321]" {
		t.Error("Expected", "[77, 22, 77, 123, 321]", ", got ", myQueue.String())
	}

	myQueue.enqueue(-15)
	myQueue.enqueue(1000)

	if myQueue.String() != "[77, 22, 77, 123, 321, -15, 1000]" {
		t.Error("Expected", "[77, 22, 77, 123, 321, -15, 1000]", ", got ", myQueue.String())
	}
}

func TestQueuePop(t *testing.T) {
	myQueue := NewQueue()
	topValue, ok := myQueue.dequeue()

	if topValue != 0 && ok != false {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[]" {
		t.Error("Expected", "[]", ", got ", myQueue.String())
	}

	myQueue = NewQueue(0, 5, 10, 15, 20, 25, 777, 32)
	topValue, ok = myQueue.dequeue()

	if topValue != 0 && ok != true {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[5, 10, 15, 20, 25, 777, 32]" {
		t.Error("Expected", "[5, 10, 15, 20, 25, 777, 32]", ", got ", myQueue.String())
	}

	topValue, ok = myQueue.dequeue()

	if topValue != 5 && ok != true {
		t.Error("Expected", 5, ", got ", topValue)
	}

	if myQueue.String() != "[10, 15, 20, 25, 777, 32]" {
		t.Error("Expected", "[10, 15, 20, 25, 777, 32]", ", got ", myQueue.String())
	}
	topValue, ok = myQueue.dequeue()
	topValue, ok = myQueue.dequeue()
	topValue, ok = myQueue.dequeue()
	topValue, ok = myQueue.dequeue()

	if topValue != 25 && ok != true {
		t.Error("Expected", 25, ", got ", topValue)
	}

	if myQueue.String() != "[777, 32]" {
		t.Error("Expected", "[777, 32]", ", got ", myQueue.String())
	}
}

func TestQueuePeek(t *testing.T) {
	myQueue := NewQueue()
	topValue, ok := myQueue.peek()

	if topValue != 0 && ok != false {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[]" {
		t.Error("Expected", "[]", ", got ", myQueue.String())
	}

	myQueue = NewQueue(0, 5, 10, 15, 20, 25, 777, 32)
	topValue, ok = myQueue.peek()

	if topValue != 0 && ok != true {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[0, 5, 10, 15, 20, 25, 777, 32]" {
		t.Error("Expected", "[0, 5, 10, 15, 20, 25, 777, 32]", ", got ", myQueue.String())
	}

	topValue, ok = myQueue.peek()
	topValue, ok = myQueue.dequeue()
	topValue, ok = myQueue.peek()

	if topValue != 0 && ok != true {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[5, 10, 15, 20, 25, 777, 32]" {
		t.Error("Expected", "[5, 10, 15, 20, 25, 777, 32]", ", got ", myQueue.String())
	}
}
