package bst

import (
	"testing"
	"reflect"
)

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

func TestBinarySearchTree_IsBalanced(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Add(9)
	bst.Add(4)
	bst.Add(17)
	bst.Add(3)
	bst.Add(6)
	bst.Add(22)
	bst.Add(5)
	bst.Add(7)
	bst.Add(20)

	min := bst.FindMinHeight()
	assetEqual(t, 1, min)

	max := bst.FindMaxHeight()
	assetEqual(t, 3, max)

	balanced := bst.IsBalanced()
	assetEqual(t, false, balanced)

	bst.Add(10)

	min = bst.FindMinHeight()
	assetEqual(t, 2, min)

	max = bst.FindMaxHeight()
	assetEqual(t, 3, max)

	balanced = bst.IsBalanced()
	assetEqual(t, true, balanced)

}

func TestBinarySearchTree_InOrderTraversal(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Add(9)
	bst.Add(4)
	bst.Add(17)
	bst.Add(3)
	bst.Add(6)
	bst.Add(22)
	bst.Add(5)
	bst.Add(7)
	bst.Add(20)
	bst.Add(10)

	list := bst.InOrderTraversal()
	expectedList := []int{3, 4, 5, 6, 7, 9, 10, 17, 20, 22}

	if !reflect.DeepEqual(list, expectedList) {
		t.Error("Expected", expectedList, ", got ", list)
	}

}

func TestBinarySearchTree_PreOrderTraversal(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Add(9)
	bst.Add(4)
	bst.Add(17)
	bst.Add(3)
	bst.Add(6)
	bst.Add(22)
	bst.Add(5)
	bst.Add(7)
	bst.Add(20)
	bst.Add(10)

	list := bst.PreOrderTraversal()
	expectedList := []int{9, 4, 3, 6, 5, 7, 17, 10, 22, 20}

	if !reflect.DeepEqual(list, expectedList) {
		t.Error("Expected", expectedList, ", got ", list)
	}

}

func TestBinarySearchTree_PostOrderTraversal(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Add(9)
	bst.Add(4)
	bst.Add(17)
	bst.Add(3)
	bst.Add(6)
	bst.Add(22)
	bst.Add(5)
	bst.Add(7)
	bst.Add(20)
	bst.Add(10)

	list := bst.PostOrderTraversal()
	expectedList := []int{3, 5, 7, 6, 4, 10, 20, 22, 17, 9}

	if !reflect.DeepEqual(list, expectedList) {
		t.Error("Expected", expectedList, ", got ", list)
	}
}
