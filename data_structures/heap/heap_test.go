package heap

import (
	"testing"
)

func assetEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected `%v`. Got `%v`\n", expected, actual)
	}
}

func TestNewHashMap(t *testing.T) {
	h := NewHeap()
	h.Push(3)
	h.Push(10)
	h.Push(1)
	h.Push(55)
	h.Push(13)
	h.Push(500)
	h.Push(77)
	h.Push(30)
	h.Push(5)

	assetEqual(t, 500, h.Pop())
	assetEqual(t, 77, h.Pop())
	assetEqual(t, 55, h.Pop())
	assetEqual(t, 30, h.Pop())
	assetEqual(t, 13, h.Pop())
	assetEqual(t, 10, h.Pop())
	assetEqual(t, 5, h.Pop())
	assetEqual(t, 3, h.Pop())
	assetEqual(t, 1, h.Pop())
}
