package bst

import "testing"

func assetEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()

	if expected != actual {
		t.Errorf("Expected `%v`. Got `%v`\n", expected, actual)
	}
}

func TestBinarySearchTreeMin(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Add(36)
	min := bst.Min()
	assetEqual(t, 36, min)

	bst.Add(10)
	min = bst.Min()
	assetEqual(t, 10, min)

	bst.Add(15)
	min = bst.Min()
	assetEqual(t, 10, min)

	bst.Add(25)
	min = bst.Min()
	assetEqual(t, 10, min)


	bst.Add(5)
	min = bst.Min()
	assetEqual(t, 5, min)

	bst.Add(1)
	min = bst.Min()
	assetEqual(t, 1, min)

	bst.Add(100)
	min = bst.Min()
	assetEqual(t, 1, min)
}

func TestBinarySearchTreeMax(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Add(25)
	max := bst.Max()
	assetEqual(t, 25, max)

	bst.Add(10)
	max = bst.Max()
	assetEqual(t, 25, max)

	bst.Add(15)
	max = bst.Max()
	assetEqual(t, 25, max)

	bst.Add(36)
	max = bst.Max()
	assetEqual(t, 36, max)


	bst.Add(5)
	max = bst.Max()
	assetEqual(t, 36, max)

	bst.Add(1)
	max = bst.Max()
	assetEqual(t, 36, max)

	bst.Add(100)
	max = bst.Max()
	assetEqual(t, 100, max)
}


func TestBinarySearchTreeFind(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Add(25)
	bst.Add(10)
	bst.Add(15)
	bst.Add(36)
	bst.Add(5)
	bst.Add(100)
	bst.Add(1)

	val := bst.Find(25)
	assetEqual(t, 25, val.value)

	val = bst.Find(75)
	if val != nil {
		t.Errorf("Expected `%v`. Got `%v`\n", nil, val)
	}

	val = bst.Find(100)
	assetEqual(t, 100, val.value)

	val = bst.Find(99)
	if val != nil {
		t.Errorf("Expected `%v`. Got `%v`\n", nil, val)
	}


	val = bst.Find(5)
	assetEqual(t, 5, val.value)

	val = bst.Find(3)
	if val != nil {
		t.Errorf("Expected `%v`. Got `%v`\n", nil, val)
	}

	bst.Add(99)
	val = bst.Find(99)
	assetEqual(t, 99, val.value)
}



func TestBinarySearchTreeIsPresent(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Add(25)
	bst.Add(10)
	bst.Add(15)
	bst.Add(36)
	bst.Add(5)
	bst.Add(100)
	bst.Add(1)

	val := bst.IsPresent(25)
	assetEqual(t, true, val)

	val = bst.IsPresent(75)
	assetEqual(t, false, val)

	val = bst.IsPresent(100)
	assetEqual(t, true, val)

	val = bst.IsPresent(99)
	assetEqual(t, false, val)


	val = bst.IsPresent(5)
	assetEqual(t, true, val)

	val = bst.IsPresent(3)
	assetEqual(t, false, val)

	bst.Add(99)
	val = bst.IsPresent(99)
	assetEqual(t, true, val)
}



func TestBinarySearchTreeRemove(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Add(25)
	bst.Add(10)
	bst.Add(15)
	bst.Add(36)
	bst.Add(5)
	bst.Add(100)
	bst.Add(1)

	val := bst.IsPresent(25)
	assetEqual(t, true, val)
	bst.Remove(25)
	val = bst.IsPresent(25)
	assetEqual(t, false, val)

	val = bst.IsPresent(36)
	assetEqual(t, true, val)
	bst.Remove(36)
	val = bst.IsPresent(36)
	assetEqual(t, false, val)

	val = bst.IsPresent(100)
	assetEqual(t, true, val)
	bst.Remove(100)
	val = bst.IsPresent(100)
	assetEqual(t, false, val)

	val = bst.IsPresent(1)
	assetEqual(t, true, val)
	bst.Remove(1)
	val = bst.IsPresent(1)
	assetEqual(t, false, val)

}

