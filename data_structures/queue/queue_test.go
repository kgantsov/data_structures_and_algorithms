package queue

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

func TestQueue_Enqueue(t *testing.T) {
	myQueue := NewQueue()
	myQueue.Enqueue(10)
	myQueue.Enqueue(15)
	myQueue.Enqueue(90)
	myQueue.Enqueue(35)

	if myQueue.String() != "[10, 15, 90, 35]" {
		t.Error("Expected", "[10, 15, 90, 35]", ", got ", myQueue.String())
	}

	myQueue = NewQueue()
	myQueue.Enqueue(77)

	if myQueue.String() != "[77]" {
		t.Error("Expected", "[77]", ", got ", myQueue.String())
	}
	myQueue.Enqueue(22)
	myQueue.Enqueue(77)
	myQueue.Enqueue(123)
	myQueue.Enqueue(321)

	if myQueue.String() != "[77, 22, 77, 123, 321]" {
		t.Error("Expected", "[77, 22, 77, 123, 321]", ", got ", myQueue.String())
	}

	myQueue.Enqueue(-15)
	myQueue.Enqueue(1000)

	if myQueue.String() != "[77, 22, 77, 123, 321, -15, 1000]" {
		t.Error("Expected", "[77, 22, 77, 123, 321, -15, 1000]", ", got ", myQueue.String())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	myQueue := NewQueue()
	topValue, ok := myQueue.Dequeue()

	if topValue != 0 && ok != false {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[]" {
		t.Error("Expected", "[]", ", got ", myQueue.String())
	}

	myQueue = NewQueue(0, 5, 10, 15, 20, 25, 777, 32)
	topValue, ok = myQueue.Dequeue()

	if topValue != 0 && ok != true {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[5, 10, 15, 20, 25, 777, 32]" {
		t.Error("Expected", "[5, 10, 15, 20, 25, 777, 32]", ", got ", myQueue.String())
	}

	topValue, ok = myQueue.Dequeue()

	if topValue != 5 && ok != true {
		t.Error("Expected", 5, ", got ", topValue)
	}

	if myQueue.String() != "[10, 15, 20, 25, 777, 32]" {
		t.Error("Expected", "[10, 15, 20, 25, 777, 32]", ", got ", myQueue.String())
	}
	topValue, ok = myQueue.Dequeue()
	topValue, ok = myQueue.Dequeue()
	topValue, ok = myQueue.Dequeue()
	topValue, ok = myQueue.Dequeue()

	if topValue != 25 && ok != true {
		t.Error("Expected", 25, ", got ", topValue)
	}

	if myQueue.String() != "[777, 32]" {
		t.Error("Expected", "[777, 32]", ", got ", myQueue.String())
	}
}

func TestQueue_Peek(t *testing.T) {
	myQueue := NewQueue()
	topValue, ok := myQueue.Peek()

	if topValue != 0 && ok != false {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[]" {
		t.Error("Expected", "[]", ", got ", myQueue.String())
	}

	myQueue = NewQueue(0, 5, 10, 15, 20, 25, 777, 32)
	topValue, ok = myQueue.Peek()

	if topValue != 0 && ok != true {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[0, 5, 10, 15, 20, 25, 777, 32]" {
		t.Error("Expected", "[0, 5, 10, 15, 20, 25, 777, 32]", ", got ", myQueue.String())
	}

	topValue, ok = myQueue.Peek()
	topValue, ok = myQueue.Dequeue()
	topValue, ok = myQueue.Peek()

	if topValue != 0 && ok != true {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[5, 10, 15, 20, 25, 777, 32]" {
		t.Error("Expected", "[5, 10, 15, 20, 25, 777, 32]", ", got ", myQueue.String())
	}
}


func TestQueue_Empty(t *testing.T) {
	myQueue := NewQueue()
	topValue, ok := myQueue.Dequeue()

	if topValue != 0 && ok != false {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.String() != "[]" {
		t.Error("Expected", "[]", ", got ", myQueue.String())
	}

	if myQueue.Empty() != true {
		t.Error("Expected", true, ", got ", myQueue.Empty())
	}

	myQueue = NewQueue(0, 5, 10)

	if myQueue.Empty() != false {
		t.Error("Expected", false, ", got ", myQueue.Empty())
	}

	topValue, ok = myQueue.Dequeue()

	if topValue != 0 && ok != true {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.Empty() != false {
		t.Error("Expected", false, ", got ", myQueue.Empty())
	}

	topValue, ok = myQueue.Dequeue()

	if topValue != 0 && ok != true {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.Empty() != false {
		t.Error("Expected", false, ", got ", myQueue.Empty())
	}

	topValue, ok = myQueue.Dequeue()

	if topValue != 0 && ok != true {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myQueue.Empty() != true {
		t.Error("Expected", true, ", got ", myQueue.Empty())
	}

}