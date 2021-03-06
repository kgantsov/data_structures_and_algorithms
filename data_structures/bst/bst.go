package bst

import "github.com/kgantsov/data_structures_and_algorithms/data_structures/queue"

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	node := new(Node)
	node.value = value

	return node
}

type BinarySearchTree struct {
	root  *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	bst := new(BinarySearchTree)
	bst.root = nil

	return bst
}

func addNode(value int, node *Node) *Node {
	if value < node.value {
		if node.left == nil {
			node.left = NewNode(value)
			return nil
		} else {
			return addNode(value, node.left)
		}
	} else if value > node.value {
		if value > node.value {
			if node.right == nil {
				node.right = NewNode(value)
				return nil
			} else {
				return addNode(value, node.right)
			}
		}
	} else {
		return nil
	}
	return nil
}

func removeNode(value int, node *Node) *Node {
	if node == nil {
		return nil
	}

	if value == node.value {
		if node.left == nil && node.right == nil {
			return nil
		}

		if node.left == nil {
			return node.right
		}

		if node.right == nil {
			return node.left
		}

		tempNode := node.right

		for tempNode.left != nil {
			tempNode = tempNode.left
		}
		node.value = tempNode.value
		node.right = removeNode(node.value, node.right)
		return node
	} else if value < node.value {
		node.left = removeNode(value, node.left)
		return node
	} else {
		node.right = removeNode(value, node.right)
		return node
	}
	return nil
}

func findMinHeight(node *Node) int {
	if node == nil {
		return -1
	}
	left := findMinHeight(node.left)
	right := findMinHeight(node.right)

	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}

func findMaxHeight(node *Node) int {
	if node == nil {
		return -1
	}
	left := findMaxHeight(node.left)
	right := findMaxHeight(node.right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func inOrderTraversal(node *Node) []int {
	var res []int

	if node.left != nil {
		res = append(res, inOrderTraversal(node.left)...)
	}

	res = append(res, node.value)

	if node.right != nil {
		res = append(res, inOrderTraversal(node.right)...)
	}
	return res
}

func preOrderTraversal(node *Node) []int {
	var res []int

	res = append(res, node.value)

	if node.left != nil {
		res = append(res, preOrderTraversal(node.left)...)
	}

	if node.right != nil {
		res = append(res, preOrderTraversal(node.right)...)
	}
	return res
}

func postOrderTraversal(node *Node) []int {
	var res []int

	if node.left != nil {
		res = append(res, postOrderTraversal(node.left)...)
	}

	if node.right != nil {
		res = append(res, postOrderTraversal(node.right)...)
	}

	res = append(res, node.value)

	return res
}


func (bst *BinarySearchTree) Add(value int) {
	node := bst.root
	if node == nil {
		bst.root = NewNode(value)
		return
	} else {
		addNode(value, node)
		return
	}
}

func (bst *BinarySearchTree) Min() int {
	node := bst.root

	if node == nil {
		return 0
	}

	for node.left != nil {
		node = node.left
	}
	return node.value
}

func (bst *BinarySearchTree) Max() int {
	node := bst.root

	if node == nil {
		return 0
	}

	for node.right != nil {
		node = node.right
	}
	return node.value
}

func (bst *BinarySearchTree) Find(value int) *Node {
	node := bst.root

	if node == nil {
		return nil
	}

	for value != node.value {
		if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
		if node == nil {
			return nil
		}
	}
	return node
}

func (bst *BinarySearchTree) IsPresent(value int) bool {
	node := bst.root

	if node == nil {
		return false
	}

	for value != node.value {
		if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
		if node == nil {
			return false
		}
	}
	return true
}

func (bst *BinarySearchTree) Remove(value int) {
	bst.root = removeNode(value, bst.root)
}

func (bst *BinarySearchTree) FindMinHeight() int {
	return findMinHeight(bst.root)
}

func (bst *BinarySearchTree) FindMaxHeight() int {
	return findMaxHeight(bst.root)
}

func (bst *BinarySearchTree) IsBalanced() bool {
	min := bst.FindMinHeight()
	max := bst.FindMaxHeight()

	if min >= max - 1 {
		return true
	} else {
		return false
	}
}

func (bst *BinarySearchTree) InOrderTraversal() []int {
	values := inOrderTraversal(bst.root)

	return values
}

func (bst *BinarySearchTree) PreOrderTraversal() []int {
	values := preOrderTraversal(bst.root)

	return values
}

func (bst *BinarySearchTree) PostOrderTraversal() []int {
	values := postOrderTraversal(bst.root)

	return values
}

func (bst *BinarySearchTree) LevelOrderTraversal() []int {
	var values []int
	var node *Node
	q := queue.NewQueue()

	if bst.root != nil {
		q.Enqueue(bst.root)

		for q.Empty() != true {
			v, ok := q.Dequeue()
			if ok {
				node = v.(*Node)

				values = append(values, node.value)
				if node.left != nil {
					q.Enqueue(node.left)
				}
				if node.right != nil {
					q.Enqueue(node.right)
				}
			}
		}
	}

	return values
}
