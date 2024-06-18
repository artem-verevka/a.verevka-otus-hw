package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.first = &ListItem{Value: v, Next: l.first}

	if l.first.Next != nil {
		l.first.Next.Prev = l.first
	}

	if l.last == nil {
		l.last = l.first
	}

	l.len++
	return l.first
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.last = &ListItem{Value: v, Prev: l.last}

	if l.last.Prev != nil {
		l.last.Prev.Next = l.last
	}

	if l.first == nil {
		l.first = l.last
	}

	l.len++
	return l.last
}

func (l *list) Remove(item *ListItem) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
	} else {
		l.first = item.Next
	}

	if item.Next != nil {
		item.Next.Prev = item.Prev
	} else {
		l.last = item.Prev
	}

	l.len--
}

func (l *list) MoveToFront(item *ListItem) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
		if item.Next != nil {
			item.Next.Prev = item.Prev
		} else {
			l.last = item.Prev
		}
		l.first.Prev = item
		item.Next = l.first
		item.Prev = nil
		l.first = item
	}
}
