package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(interface{}) *listItem
	PushBack(interface{}) *listItem
	Remove(item *listItem)
	MoveToFront(item *listItem)
}

type listItem struct {
	Value interface{}
	Next *listItem
	Previous *listItem
}

type list struct {
	head *listItem
	tail *listItem
	size int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *listItem {
	return l.head
}

func (l *list) Back() *listItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *listItem {
	n := &listItem{Value: v}
	if l.size == 0 {
		l.head = n
		l.tail = n
		l.size++
		return n
	}
	firstItem := l.head
	l.head = n
	l.head.Next = firstItem
	firstItem.Previous = n
	l.size++
	return n
}

func (l *list) PushBack(v interface{}) *listItem {
	n := &listItem{Value: v}
	if l.size == 0 {
		l.head = n
		l.tail = n
		l.size++
		return n
	}
	lastItem := l.tail
	l.tail = n
	l.tail.Previous = lastItem
	lastItem.Next = n
	l.size++
	return n
}

func (l *list) Remove(v *listItem) {
	if l.size == 0 {
		return
	}

	if l.head.Value == v.Value {
		ln := l.head.Next
		if ln != nil {
			ln.Previous = nil
			l.head = ln
		} else {
			l.head = nil
			l.tail = nil
		}
		l.size--
		return
	}

	if l.tail.Value == v.Value {
		ln := l.tail.Previous
		l.tail = ln
		ln.Next = nil
		l.size--
		return
	}

	head := l.head
	var previousListItem = new(listItem)
	var nextListItem = new(listItem)

	for head.Value != v.Value {
		if head == nil {
			return
		}
		previousListItem = head
		head = head.Next
		nextListItem = head.Next
	}

	previousListItem.Next = head.Next
	nextListItem.Previous = previousListItem
	l.size--
}

func (l *list) MoveToFront(v *listItem) {
	l.Remove(v)
	l.PushFront(v.Value)
}