package linklist

import "testing"

// test the Link reverse() interface using list of size n
func testReverse(n int) bool {

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

func testLength(n int) bool {

	list := Link{0, nil}
	for i := 1; i <= n-1; i++ {
		list.append(i)
	}
	if list.length() == n {
		return true
	}
	return false
}

func TestLength(t *testing.T) {

	cases := []int{1, 2, 10}
	for _, n := range cases {
		if !testLength(n) {
			t.Errorf("Fail %d", n)
		}
	}
}

func TestReverse(t *testing.T) {

	cases := []int{1, 2, 10}
	for _, n := range cases {
		if !testReverse(n) {
			t.Errorf("Fail %d", n)
		}
	}
}

func TestAppend(t *testing.T) {

	l := Link{0, nil}
	l.append(1)
	if l.length() != 2 {
		t.Errorf("Length Incorrect")
	} else if l.next.value != 1 {
		t.Errorf("Value incorret")
	}
}
