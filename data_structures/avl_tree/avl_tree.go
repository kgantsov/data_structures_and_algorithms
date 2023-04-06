package avl_tree

type Node struct {
	value  int
	height int
	left   *Node
	right  *Node
}

func NewNode(value int) *Node {
	node := &Node{
		value:  value,
		height: 1,
		left:   nil,
		right:  nil,
	}

	return node
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *Node) balanceFactor() int {
	if n == nil {
		return 0
	}

	return n.left.Height() - n.right.Height()
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (n *Node) updateHeight() {
	n.height = max(n.left.Height(), n.right.Height()) + 1
}

func rightRotate(x *Node) *Node {
	y := x.left
	t := y.right

	y.right = x
	x.left = t

	x.updateHeight()
	y.updateHeight()

	return y
}

func leftRotate(x *Node) *Node {
	y := x.right
	t := y.left

	y.left = x
	x.right = t

	x.updateHeight()
	y.updateHeight()

	return y
}

type AVLTree struct {
	root *Node
}

func NewAVLTree() *AVLTree {
	bst := &AVLTree{root: nil}
	return bst
}

func (t *AVLTree) Insert(value int) {
	t.root = t.insert(t.root, value)
}

func (t *AVLTree) insert(node *Node, value int) *Node {
	// Step 1 - Perform normal BST
	if node == nil {
		return NewNode(value)
	} else if value < node.value {
		node.left = t.insert(node.left, value)
	} else {
		node.right = t.insert(node.right, value)
	}

	// Step 2 - Update the height of the ancestor node
	node.updateHeight()

	// Step 3 - Get the balance factor
	balance := node.balanceFactor()

	// Step 4 - If the node is unbalanced,
	// then try out the 4 cases
	// Case 1 - Left Left
	if balance > 1 && value < node.left.value {
		return rightRotate(node)
	}

	// Case 2 - Right Right
	if balance < -1 && value > node.right.value {
		return leftRotate(node)
	}

	// Case 3 - Left Right
	if balance > 1 && value > node.left.value {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// Case 4 - Right Left
	if balance < -1 && value < node.right.value {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func (t *AVLTree) getMinValueNode(node *Node) *Node {
	if node == nil || node.left == nil {
		return node
	}

	return t.getMinValueNode(node.left)
}

func (t *AVLTree) Delete(value int) {
	t.root = t.delete(t.root, value)
}

func (t *AVLTree) delete(node *Node, value int) *Node {
	// Step 1 - Perform normal BST
	if node == nil {
		return nil
	} else if value < node.value {
		node.left = t.delete(node.left, value)
	} else if value > node.value {
		node.right = t.delete(node.right, value)
	} else {
		if node.left == nil {
			temp := node.right
			node = nil
			return temp
		} else if node.right == nil {
			temp := node.left
			node = nil
			return temp
		}

		temp := t.getMinValueNode(node.right)
		node.value = temp.value
		node.right = t.delete(node.right, temp.value)
	}

	if node == nil {
		return nil
	}

	// Step 2 - Update the height of the ancestor node
	node.updateHeight()

	// Step 3 - Get the balance factor
	balance := node.balanceFactor()

	// Step 4 - If the node is unbalanced,
	// then try out the 4 cases
	// Case 1 - Left Left
	if balance > 1 && node.left.balanceFactor() >= 0 {
		return rightRotate(node)
	}

	// Case 2 - Right Right
	if balance < -1 && node.right.balanceFactor() <= 0 {
		return leftRotate(node)
	}

	// Case 3 - Left Right
	if balance > 1 && node.left.balanceFactor() < 0 {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// Case 4 - Right Left
	if balance < -1 && node.right.balanceFactor() > 0 {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func (t *AVLTree) InOrderTraversal() []int {
	values := t.inOrderTraversal(t.root)

	return values
}

func (t *AVLTree) inOrderTraversal(node *Node) []int {
	res := []int{}

	if node == nil {
		return res
	}

	if node.left != nil {
		res = append(res, t.inOrderTraversal(node.left)...)
	}

	res = append(res, node.value)

	if node.right != nil {
		res = append(res, t.inOrderTraversal(node.right)...)
	}

	return res
}

func (t *AVLTree) Height() int {
	return t.root.Height()
}
