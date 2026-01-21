package pkg

/*
The doubly linked list is a foundational data structure with O(1) / constant time complexity to access the head or the
tail of the list and O(n) / linear time complexity to iterate
*/

type DoublyLinkedListItem[T any] struct {
	value T
	next  *DoublyLinkedListItem[T]
	prev  *DoublyLinkedListItem[T]
}

type DoublyLinkedList[T any] struct {
	head   *DoublyLinkedListItem[T]
	tail   *DoublyLinkedListItem[T]
	length int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	l := &DoublyLinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}

	return l
}

func (l *DoublyLinkedList[T]) Push(value T) {
	item := &DoublyLinkedListItem[T]{
		value: value,
		next:  nil,
		prev:  l.tail,
	}

	if l.head == nil {
		l.head = item
	}

	if l.tail == nil {
		l.tail = item
	} else {
		l.tail.next = item

		l.tail = item
	}

	l.length += 1
}

func (l *DoublyLinkedList[T]) Pop() *T {
	t := l.tail

	if t == nil {
		return nil
	}

	if l.head == t {
		l.head = nil
	}

	if t.prev != nil {
		t.prev.next = nil
	}

	l.tail = t.prev

	l.length -= 1

	return &t.value
}

func (l *DoublyLinkedList[T]) PushLeft(value T) {
	item := &DoublyLinkedListItem[T]{
		value: value,
		next:  l.head,
		prev:  nil,
	}

	if l.head == nil {
		l.head = item
	} else {
		l.head.prev = item
		l.head = item
	}

	if l.tail == nil {
		l.tail = item
	}

	l.length += 1
}

func (l *DoublyLinkedList[T]) PopLeft() *T {
	t := l.head

	if t == nil {
		return nil
	}

	if l.tail == t {
		l.tail = nil
	}

	if t.next != nil {
		t.next.prev = nil
	}

	l.head = t.next

	l.length -= 1

	return &t.value
}

func (l *DoublyLinkedList[T]) Length() int {
	return l.length
}

func (l *DoublyLinkedList[T]) Items() []T {
	items := make([]T, 0)

	n := l.head

	for n != nil {
		items = append(items, n.value)

		n = n.next
	}

	return items
}
