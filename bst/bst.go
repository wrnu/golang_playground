package bst

// TreeNode ... Node Type for a BST
type TreeNode struct {
	key   int64
	value int64
	left  *TreeNode
	right *TreeNode
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

func (node *TreeNode) isBST(minKey int64, maxKey int64) bool {

	if node == nil {
		return true
	}
	if (node.key < minKey) || (node.key > maxKey) {
		return false
	}
	return node.left.isBST(minKey, node.key-1) && node.right.isBST(node.key+1, maxKey)
}
