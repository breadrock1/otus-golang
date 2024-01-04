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
	length int
	front  *ListItem
	back   *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

// PushFront new element to begin of linked list.
// ItemName and fields: front, back
// new      | a,            b,      c,
// *a, nil  | *b, nil/*new  *c, *a  nil, *b
func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.length < 1 {
		l.back = newItem
	} else {
		newItem.Next = l.front
		l.front.Prev = newItem
	}

	l.front = newItem
	l.length++
	return newItem
}

// PushBack new element to end of linked list.
// ItemName and fields: front, back
// a,       b,      c,            | new
// *b, nil  *c, *a  nil/*new, *b  | nil, *c
func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.length < 1 {
		l.back = newItem
	} else {
		newItem.Prev = l.back
		l.back.Next = newItem
	}

	l.back = newItem
	l.length++
	return newItem
}

// Remove element from linked list
// ItemName and fields: front, back
// a,       b,         | c,      | d
// *b, nil  *c/*d, *a  | *d, *b  | nil, *c/*b
func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.front = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.back = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	l.Remove(i)

	i.Prev = nil
	i.Next = l.front

	l.front.Prev = i
	l.front = i

	l.length++
}

func NewList() List {
	return &list{}
}
