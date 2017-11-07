package main

import (
	"fmt"
)

type Link struct {
	value int
	next  *Link
}

func (l *Link) length() int {
	if l == nil {
		return 0
	}
	n := 0
	curr := l
	for curr != nil {
		curr = curr.next
		n++
	}
	return n
}

func (l *Link) append(value int) {
	curr := l

	if curr == nil {
		return
	}

	for curr.next != nil {
		curr = curr.next
	}

	tail := Link{value, nil}
	curr.next = &tail
}

func (l Link) reverse() *Link {
	curr := &l
	var next, prev *Link
	next = nil
	prev = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

// test the Link reverse() interface using list of size n
func test_reverse(n int) bool {

	list1 := Link{0, nil}
	for i := 1; i <= n-1; i++ {
		list1.append(i)
	}

	list2 := Link{n - 1, nil}
	for i := n - 2; i >= 0; i-- {
		list2.append(i)
	}

	curr1 := &list1
	curr2 := list2.reverse()

	for curr1 != nil && curr2 != nil {
		if curr1.value != curr2.value {
			return false
		}
		curr1 = curr1.next
		curr2 = curr2.next
	}
	return true
}

func main() {

	if test_reverse(1) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	if test_reverse(2) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	if test_reverse(10) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
