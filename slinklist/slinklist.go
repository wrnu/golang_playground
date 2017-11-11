package slinklist

// SLinkList ... Singly Linked List Type
type SLinkList struct {
	value int
	next  *SLinkList
}

func (l *SLinkList) length() int {

	n := 0
	curr := l
	for curr != nil {
		curr = curr.next
		n++
	}
	return n
}

func (l *SLinkList) prepend(value int) *SLinkList {

	head := SLinkList{value, l}
	return &head
}

func (l *SLinkList) append(value int) *SLinkList {

	curr := l

	for curr.next != nil {
		curr = curr.next
	}

	tail := SLinkList{value, nil}
	curr.next = &tail

	return l
}

func (l *SLinkList) insertAfter(value int) *SLinkList {

	new := SLinkList{value, l.next}
	l.next = &new
	return l
}

func (l *SLinkList) delete(value int) *SLinkList {
	curr := l
	if curr.value == value {
		return curr.next
	}
	for curr.next != nil {
		if curr.next.value == value {
			curr.next = curr.next.next
			return l
		}
		curr = curr.next
	}
	return l
}

func (l SLinkList) reverse() *SLinkList {

	curr := &l
	var next, prev *SLinkList
	prev = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

func (l *SLinkList) reverseRecursion() *SLinkList {

	curr := l
	if curr == nil || curr.next == nil {
		return curr
	}
	rev := curr.next.reverseRecursion()
	curr.next.next = curr
	curr.next = nil

	return rev
}

func (l *SLinkList) search(value int) *SLinkList {

	curr := l
	for curr != nil && curr.value != value {
		curr = curr.next
	}
	if curr == nil || curr.value != value {
		return nil
	}
	return curr
}
