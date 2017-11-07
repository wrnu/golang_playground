package bst

import (
	"math"
	"testing"
)

func TestInsert(t *testing.T) {
	bst := TreeNode{1, 1, nil, nil}
	bst.insert(0, 0)
	bst.insert(2, 2)
	if bst.left.key != 0 || bst.left.value != 0 {
		t.Errorf("Left Add Error")
	}
	if bst.right.key != 2 || bst.right.value != 2 {
		t.Errorf("Right Add Error")
	}
}

func TestIsBST(t *testing.T) {
	left := &TreeNode{0, 0, nil, nil}
	right := &TreeNode{2, 2, nil, nil}
	root := &TreeNode{1, 1, left, right}
	if root.isBST(math.MinInt64, math.MaxInt64) == false {
		t.Errorf("This is a BST")
	}
	root.right = left
	root.left = right
	if root.isBST(math.MinInt64, math.MaxInt64) == true {
		t.Errorf("This is not a BST")
	}
}
