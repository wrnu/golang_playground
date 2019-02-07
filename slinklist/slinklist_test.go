package slinklist

import "testing"

// test the SLinkList reverse() interface using list of size n
func testReverse(n int) bool {

  list1 := SLinkList{0, nil}
  for i := 1; i <= n-1; i++ {
    list1.append(i)
  }

  list2 := SLinkList{n - 1, nil}
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

// test the SLinkList reverse() interface using list of size n
func testReverseRecursion(n int) bool {

  list1 := SLinkList{0, nil}
  for i := 1; i <= n-1; i++ {
    list1.append(i)
  }

  list2 := SLinkList{n - 1, nil}
  for i := n - 2; i >= 0; i-- {
    list2.append(i)
  }

  curr1 := &list1
  curr2 := (&list2).reverseRecursion()

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

  list := SLinkList{0, nil}
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

func TestReverseRecursion(t *testing.T) {

  cases := []int{1, 2, 10}
  for _, n := range cases {
    if !testReverseRecursion(n) {
      t.Errorf("Fail %d", n)
    }
  }
}

func TestPrepend(t *testing.T) {

  l := SLinkList{0, nil}
  head := l.prepend(1)
  if head.length() != 2 {
    t.Errorf("Length incorrect")
  }
  if head.value != 1 {
    t.Errorf("Head value incorret")
  }
  if head.next.value != 0 {
    t.Errorf("Tail value incorret")
  }
}

func TestAppend(t *testing.T) {

  l := SLinkList{0, nil}
  l.append(1)
  if l.length() != 2 {
    t.Errorf("Length Incorrect")
  }
  if l.value != 0 {
    t.Errorf("Head value incorret")
  }
  if l.next.value != 1 {
    t.Errorf("Tail value incorret")
  }
}

func TestInsertAfter(t *testing.T) {

  l := SLinkList{0, nil}
  l.insertAfter(1)
  if l.length() != 2 {
    t.Errorf("Length Incorrect")
  }
  if l.value != 0 {
    t.Errorf("Head value incorret")
  }
  if l.next.value != 1 {
    t.Errorf("Tail value incorret")
  }
}

func TestDelete(t *testing.T) {

  l := &SLinkList{0, nil}
  l.append(-10)
  l.append(10)
  l.append(100)

  l = l.delete(-10)
  if l.search(-10) != nil && l.length() == 3 {
    t.Errorf("Failed to delete middle element")
  }

  l = l.delete(100)
  if l.search(1000) != nil && l.length() == 2 {
    t.Errorf("Failed to delete tail element")
  }

  l = l.delete(0)
  if l.search(0) != nil && l.length() == 1 {
    t.Errorf("Failed to delete head element")
  }

  len := l.length()
  l = l.delete(101)
  if l.length() != len {
    t.Errorf("Failed delete of value not in list")
  }

  l = l.delete(10)
  if l != nil {
    t.Errorf("Failed to delete last element")
  }
}

func TestSearch(t *testing.T) {

  l := SLinkList{0, nil}
  l.append(1)
  l.append(2)

  if l.search(0).value != 0 {
    t.Errorf("Failed to find node")
  }
  if l.search(1).value != 1 {
    t.Errorf("Failed to find node")
  }
  if l.search(2).value != 2 {
    t.Errorf("Failed to find node")
  }
  if l.search(3) != nil {
    t.Errorf("Found node that doesn't exist")
  }
}
