package avl_tree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTreeInsert(t *testing.T) {
	tree := NewAVLTree()
	assert.Equal(t, []int{}, tree.InOrderTraversal())

	tree.Insert(9)
	assert.Equal(t, []int{9}, tree.InOrderTraversal())

	tree.Insert(5)
	assert.Equal(t, []int{5, 9}, tree.InOrderTraversal())

	tree.Insert(10)
	assert.Equal(t, []int{5, 9, 10}, tree.InOrderTraversal())

	tree.Insert(0)
	assert.Equal(t, []int{0, 5, 9, 10}, tree.InOrderTraversal())

	tree.Insert(6)
	assert.Equal(t, []int{0, 5, 6, 9, 10}, tree.InOrderTraversal())

	tree.Insert(11)
	assert.Equal(t, []int{0, 5, 6, 9, 10, 11}, tree.InOrderTraversal())

	tree.Insert(-1)
	assert.Equal(t, []int{-1, 0, 5, 6, 9, 10, 11}, tree.InOrderTraversal())

	tree.Insert(1)
	assert.Equal(t, []int{-1, 0, 1, 5, 6, 9, 10, 11}, tree.InOrderTraversal())

	tree.Insert(2)
	assert.Equal(t, []int{-1, 0, 1, 2, 5, 6, 9, 10, 11}, tree.InOrderTraversal())
}

func TestAVLTreeDelete(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(9)
	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(0)
	tree.Insert(6)
	tree.Insert(11)
	tree.Insert(-1)
	tree.Insert(1)
	tree.Insert(2)

	assert.Equal(t, []int{-1, 0, 1, 2, 5, 6, 9, 10, 11}, tree.InOrderTraversal())

	tree.Delete(6)
	assert.Equal(t, []int{-1, 0, 1, 2, 5, 9, 10, 11}, tree.InOrderTraversal())

	tree.Delete(1)
	assert.Equal(t, []int{-1, 0, 2, 5, 9, 10, 11}, tree.InOrderTraversal())

	tree.Delete(11)
	assert.Equal(t, []int{-1, 0, 2, 5, 9, 10}, tree.InOrderTraversal())

	tree.Delete(-1)
	assert.Equal(t, []int{0, 2, 5, 9, 10}, tree.InOrderTraversal())

	tree.Delete(9)
	assert.Equal(t, []int{0, 2, 5, 10}, tree.InOrderTraversal())

	tree.Delete(10)
	assert.Equal(t, []int{0, 2, 5}, tree.InOrderTraversal())

	tree.Delete(0)
	assert.Equal(t, []int{2, 5}, tree.InOrderTraversal())

	tree.Delete(5)
	assert.Equal(t, []int{2}, tree.InOrderTraversal())

	tree.Delete(2)
	assert.Equal(t, []int{}, tree.InOrderTraversal())
}

func TestAVLTreeHeight(t *testing.T) {
	tree := NewAVLTree()
	items := []int{}

	for i := 1; i < int(math.Pow(2, 10)); i++ {
		items = append(items, i)
		tree.Insert(i)
	}
	assert.Equal(t, items, tree.InOrderTraversal())
	assert.Equal(t, 10, tree.Height())

	tree = NewAVLTree()
	items = []int{}

	for i := 1; i < int(math.Pow(2, 20)); i++ {
		items = append(items, i)
		tree.Insert(i)
	}
	assert.Equal(t, items, tree.InOrderTraversal())
	assert.Equal(t, 20, tree.Height())
}
