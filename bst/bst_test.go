package bst

import (
	"math"
	"testing"
)

func TestFindMin(t *testing.T) {
	bst := TreeNode{1, 1, nil, nil}
	bst.insert(0, 0)
	bst.insert(2, 2)
	if bst.findMin() != bst.left {
		t.Errorf("Wrong Min")
	}
}

func TestFindMax(t *testing.T) {
	bst := TreeNode{1, 1, nil, nil}
	bst.insert(0, 0)
	bst.insert(2, 2)
	if bst.findMax() != bst.right {
		t.Errorf("Wrong Max")
	}
}

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

func TestSearch(t *testing.T) {
	left := &TreeNode{0, 0, nil, nil}
	right := &TreeNode{2, 2, nil, nil}
	root := &TreeNode{1, 1, left, right}
	if root.search(0) != left {
		t.Errorf("Incorrect Left Node")
	}
	if root.search(2) != right {
		t.Errorf("Incorrect Right Node")
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