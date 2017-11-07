package bst

// TreeNode ... Node Type for a BST
type TreeNode struct {
	key   int64
	value int64
	left  *TreeNode
	right *TreeNode
}

func (node *TreeNode) findMin() *TreeNode {
	curr := node
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (node *TreeNode) findMax() *TreeNode {
	curr := node
	for curr.right != nil {
		curr = curr.right
	}
	return curr
}

func (node *TreeNode) delete(key int64) *TreeNode {

	if key < node.key {
		node.left = node.left.delete(key)
	} else if key > node.key {
		node.right = node.right.delete(key)
	} else {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		temp := node.right.findMin()
		node.key = temp.key
		node.right = node.right.delete(temp.key)
	}
	return node
}

func (node *TreeNode) insert(key int64, value int64) *TreeNode {
	if node == nil {
		node = &TreeNode{key, value, nil, nil}
	} else if key < node.key {
		node.left = node.left.insert(key, value)
	} else {
		node.right = node.right.insert(key, value)
	}
	return node
}

func (node *TreeNode) search(key int64) *TreeNode {
	if node == nil || node.key == key {
		return node
	}
	if key < node.key {
		return node.left.search(key)
	}
	return node.right.search(key)
}

func (node *TreeNode) isBST(minKey int64, maxKey int64) bool {

	if node == nil {
		return true
	}
	if (node.key < minKey) || (node.key > maxKey) {
		return false
	}
	return node.left.isBST(minKey, node.key-1) && node.right.isBST(node.key+1, maxKey)
}
