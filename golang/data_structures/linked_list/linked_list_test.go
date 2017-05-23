package linked_list

import (
	"testing"
)

func TestLinkedListAppend(t *testing.T) {
	list := NewLinkedList()
	list.append(10)

	if list.String() != "[10]" {
		t.Error("Expected", "[10]", ", got ", list.String())
	}

	if list.head.value != 10 {
		t.Error("Expected", 10, ", got ", list.head.value)
	}

	if list.tail.value != 10 {
		t.Error("Expected", 10, ", got ", list.tail.value)
	}
	list.append(255)

	if list.String() != "[10, 255]" {
		t.Error("Expected", "[10, 255]", ", got ", list.String())
	}

	if list.head.value != 10 {
		t.Error("Expected", 10, ", got ", list.head.value)
	}

	if list.tail.value != 255 {
		t.Error("Expected", 255, ", got ", list.tail.value)
	}

	list = NewLinkedList()
	for i := 0; i <= 25; i += 5 {
		list.append(i)
	}

	if list.String() != "[0, 5, 10, 15, 20, 25]" {
		t.Error("Expected", "[0, 5, 10, 15, 20, 25]", ", got ", list.String())
	}
}

func TestLinkedListLen(t *testing.T) {
	list := NewLinkedList()

	if list.Len() != 0 {
		t.Error("Expected", 0, ", got ", list.Len())
	}

	for i := 0; i <= 25; i += 5 {
		list.append(i)
	}

	if list.Len() != 6 {
		t.Error("Expected", 6, ", got ", list.Len())
	}

	list = NewLinkedList()
	for i := 0; i < 100; i++ {
		list.append(i)
	}

	if list.Len() != 100 {
		t.Error("Expected", 100, ", got ", list.Len())
	}
}

func TestLinkedSearch(t *testing.T) {
	list := NewLinkedList()

	for i := 0; i <= 25; i += 5 {
		list.append(i)
	}

	value, ok := list.search(10)

	if value != 10 && ok != true {
		t.Error("Expected", 10, true, ", got ", value, ok)
	}

	value, ok = list.search(11)

	if value != 0 && ok != false {
		t.Error("Expected", 10, true, ", got ", value, ok)
	}

	value, ok = list.search(0)

	if value != 0 && ok != false {
		t.Error("Expected", 10, true, ", got ", value, ok)
	}
}

func TestLinkedIndex(t *testing.T) {
	list := NewLinkedList()

	for i := 0; i <= 25; i += 5 {
		list.append(i)
	}

	index, ok := list.index(10)

	if index != 2 && ok != true {
		t.Error("Expected", 2, true, ", got ", index, ok)
	}

	index, ok = list.index(25)

	if index != 5 && ok != true {
		t.Error("Expected", 5, true, ", got ", index, ok)
	}

	index, ok = list.index(11)

	if index != 0 && ok != false {
		t.Error("Expected", 10, true, ", got ", index, ok)
	}

	index, ok = list.index(0)

	if index != 0 && ok != false {
		t.Error("Expected", 10, true, ", got ", index, ok)
	}
}
