package linklist

// Link ... Singly Linked List Type
type Link struct {
	value int
	next  *Link
}

func (l *Link) length() int {

	n := 0
	curr := l
	for curr != nil {
		curr = curr.next
		n++
	}
	return n
}
func (l *Link) prepend(value int) *Link {

	head := Link{value, l}
	return &head
}

func (l *Link) append(value int) *Link {

	curr := l

	for curr.next != nil {
		curr = curr.next
	}

	tail := Link{value, nil}
	curr.next = &tail

	return l
}

func (l *Link) insertAfter(value int) *Link {

	new := Link{value, l.next}
	l.next = &new
	return l
}

func (l Link) reverse() *Link {

	curr := &l
	var next, prev *Link
	prev = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

func (l *Link) search(value int) *Link {

	curr := l
	for curr != nil && curr.value != value {
		curr = curr.next
	}
	if curr == nil || curr.value != value {
		return nil
	}
	return curr
}
