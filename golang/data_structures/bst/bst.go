package bst

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
