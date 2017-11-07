package linklist

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
