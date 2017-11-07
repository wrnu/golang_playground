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

func TestDelete(t *testing.T) {
	left := &TreeNode{0, 0, nil, nil}
	right := &TreeNode{2, 2, nil, nil}
	root := &TreeNode{1, 1, left, right}

	root = root.delete(0)
	if root.left != nil && root.isBST(math.MinInt64, math.MaxInt64) {
		t.Errorf("Left Delete Error")
	}
	left = &TreeNode{0, 0, nil, nil}
	right = &TreeNode{2, 2, nil, nil}
	root = &TreeNode{1, 1, left, right}

	root = root.delete(2)
	if root.right != nil && root.isBST(math.MinInt64, math.MaxInt64) {
		t.Errorf("Right Delete Error")
	}

	left = &TreeNode{0, 0, nil, nil}
	right = &TreeNode{2, 2, nil, nil}
	root = &TreeNode{1, 1, left, right}

	root = root.delete(1)
	if root.key != 2 && root.isBST(math.MinInt64, math.MaxInt64) {
		t.Errorf("Root Delete Error with left and right leaves")
	}

	left = &TreeNode{0, 0, nil, nil}
	root = &TreeNode{1, 1, left, nil}

	root = root.delete(1)
	if root.key != 0 && root.isBST(math.MinInt64, math.MaxInt64) {
		t.Errorf("Root Delete Error with left leaf only")
	}

	right = &TreeNode{2, 2, nil, nil}
	root = &TreeNode{1, 1, nil, right}

	root = root.delete(1)
	if root.key != 2 && root.isBST(math.MinInt64, math.MaxInt64) {
		t.Errorf("Root Delete Error with left leaf only")
	}

	root = &TreeNode{1, 1, nil, nil}
	root.insert(10, 10)
	root.insert(4, 4)
	root.insert(11, 11)
	root.delete(10)
	if root.search(10) != nil && root.isBST(math.MinInt64, math.MaxInt64) {
		t.Errorf("Delete two child right leaf")
	}

	root = &TreeNode{1, 1, nil, nil}
	root.insert(0, 0)
	root.insert(-10, -10)
	root.insert(-4, -4)
	root.insert(-11, -11)
	root.delete(-10)
	if root.search(-10) != nil && root.isBST(math.MinInt64, math.MaxInt64) {
		t.Errorf("Delete two child left leaf")
	}

}

func TestInsert(t *testing.T) {
	bst := TreeNode{1, 1, nil, nil}
	bst.insert(0, 0)
	bst.insert(2, 2)
	if bst.left.key != 0 || bst.left.value != 0 || !bst.isBST(math.MinInt64, math.MaxInt64) {
		t.Errorf("Left Add Error")
	}
	if bst.right.key != 2 || bst.right.value != 2 || !bst.isBST(math.MinInt64, math.MaxInt64) {
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
