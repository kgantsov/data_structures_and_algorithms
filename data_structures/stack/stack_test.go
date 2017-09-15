package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	myStack := NewStack()

	if myStack.String() != "[]" {
		t.Error("Expected", "[]", ", got ", myStack.String())
	}

	myStack = NewStack(0, 5, 10, 15, 20, 25, 777, 32)

	if myStack.String() != "[32, 777, 25, 20, 15, 10, 5, 0]" {
		t.Error("Expected", "[32, 777, 25, 20, 15, 10, 5, 0]", ", got ", myStack.String())
	}

	myStack = NewStack(0, 17, -45)

	if myStack.String() != "[-45, 17, 0]" {
		t.Error("Expected", "[-45, 17, 0]", ", got ", myStack.String())
	}
}

func TestStackPush(t *testing.T) {
	myStack := NewStack()
	myStack.push(10)
	myStack.push(15)
	myStack.push(90)
	myStack.push(35)

	if myStack.String() != "[35, 90, 15, 10]" {
		t.Error("Expected", "[35, 90, 15, 10]", ", got ", myStack.String())
	}

	myStack = NewStack()
	myStack.push(77)

	if myStack.String() != "[77]" {
		t.Error("Expected", "[77]", ", got ", myStack.String())
	}
	myStack.push(22)
	myStack.push(77)
	myStack.push(123)
	myStack.push(321)

	if myStack.String() != "[321, 123, 77, 22, 77]" {
		t.Error("Expected", "[321, 123, 77, 22, 77]", ", got ", myStack.String())
	}

	myStack.push(-15)
	myStack.push(1000)

	if myStack.String() != "[1000, -15, 321, 123, 77, 22, 77]" {
		t.Error("Expected", "[1000, -15, 321, 123, 77, 22, 77]", ", got ", myStack.String())
	}
}

func TestStackPop(t *testing.T) {
	myStack := NewStack()
	topValue, ok := myStack.pop()

	if topValue != 0 && ok != false {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myStack.String() != "[]" {
		t.Error("Expected", "[]", ", got ", myStack.String())
	}

	myStack = NewStack(0, 5, 10, 15, 20, 25, 777, 32)
	topValue, ok = myStack.pop()

	if topValue != 32 && ok != true {
		t.Error("Expected", 32, ", got ", topValue)
	}

	if myStack.String() != "[777, 25, 20, 15, 10, 5, 0]" {
		t.Error("Expected", "[777, 25, 20, 15, 10, 5, 0]", ", got ", myStack.String())
	}

	topValue, ok = myStack.pop()

	if topValue != 777 && ok != true {
		t.Error("Expected", 777, ", got ", topValue)
	}

	if myStack.String() != "[25, 20, 15, 10, 5, 0]" {
		t.Error("Expected", "[25, 20, 15, 10, 5, 0]", ", got ", myStack.String())
	}
	topValue, ok = myStack.pop()
	topValue, ok = myStack.pop()
	topValue, ok = myStack.pop()
	topValue, ok = myStack.pop()

	if topValue != 10 && ok != true {
		t.Error("Expected", 10, ", got ", topValue)
	}

	if myStack.String() != "[5, 0]" {
		t.Error("Expected", "[5, 0]", ", got ", myStack.String())
	}
}

func TestStackPeek(t *testing.T) {
	myStack := NewStack()
	topValue, ok := myStack.peek()

	if topValue != 0 && ok != false {
		t.Error("Expected", 0, ", got ", topValue)
	}

	if myStack.String() != "[]" {
		t.Error("Expected", "[]", ", got ", myStack.String())
	}

	myStack = NewStack(0, 5, 10, 15, 20, 25, 777, 32)
	topValue, ok = myStack.peek()

	if topValue != 32 && ok != true {
		t.Error("Expected", 32, ", got ", topValue)
	}

	if myStack.String() != "[32, 777, 25, 20, 15, 10, 5, 0]" {
		t.Error("Expected", "[32, 777, 25, 20, 15, 10, 5, 0]", ", got ", myStack.String())
	}

	topValue, ok = myStack.peek()
	topValue, ok = myStack.pop()
	topValue, ok = myStack.peek()

	if topValue != 777 && ok != true {
		t.Error("Expected", 777, ", got ", topValue)
	}

	if myStack.String() != "[777, 25, 20, 15, 10, 5, 0]" {
		t.Error("Expected", "[777, 25, 20, 15, 10, 5, 0]", ", got ", myStack.String())
	}
}
