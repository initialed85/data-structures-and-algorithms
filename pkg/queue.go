package pkg

/*
The queue is a FIFO data structure with O(1) / constant time complexity to push or pop and O(n) / linear time
complexity to iterate
*/

type Queue[T any] struct {
	doublyLinkedList *DoublyLinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		doublyLinkedList: NewDoublyLinkedList[T](),
	}
}

func (q *Queue[T]) Push(value T) {
	q.doublyLinkedList.Push(value)
}

func (q *Queue[T]) Pop() *T {
	return q.doublyLinkedList.PopLeft()
}

func (q *Queue[T]) Length() int {
	return q.doublyLinkedList.Length()
}

func (q *Queue[T]) Items() []T {
	return q.doublyLinkedList.Items()
}
